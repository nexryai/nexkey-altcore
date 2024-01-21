package xnote

import (
	"fmt"
	"lab.sda1.net/nexryai/altcore/internal/activitypub"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/enum"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/core/system"
	"lab.sda1.net/nexryai/altcore/internal/core/utils"
	"lab.sda1.net/nexryai/altcore/internal/db"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
	"lab.sda1.net/nexryai/altcore/internal/db/kv"
	"lab.sda1.net/nexryai/altcore/internal/queue"
	"lab.sda1.net/nexryai/altcore/internal/services"
	"lab.sda1.net/nexryai/altcore/internal/services/baselib"
)

func addToDeliverQueue(note *entities.Note) error {
	followService := baselib.FollowService{
		UserId: note.UserId,
		Type:   enum.Followers,
	}
	followers, err := followService.FindAll()
	if err != nil {
		return err
	}

	userService := baselib.UserService{
		LocalOnly: false,
	}
	inboxes, err := userService.GetSharedInboxes(followers)
	if err != nil {
		return err
	}

	logger.Warn(fmt.Sprintf("%s", inboxes))

	actorUrl := fmt.Sprintf("%s/users/%s", config.URL, note.UserId)
	id := fmt.Sprintf("%s/notes/%s/activity", config.URL, note.Id)

	for _, inbox := range inboxes {
		activity := activitypub.CreateActivity{
			Id:        id,
			Actor:     actorUrl,
			Published: note.CreatedAt,
			To:        []string{fmt.Sprintf("%s/followers", actorUrl)},
			Cc:        []string{"https://www.w3.org/ns/activitystreams#Public"},
			Object: activitypub.CreateActivityObject{
				Id:             id,
				Type:           "Note",
				Actor:          actorUrl,
				AttributedTo:   actorUrl,
				Summary:        nil,
				Content:        note.Text, // ToDo: html化
				MisskeyContent: note.Text,
				Source: activitypub.ActivityObjectSource{
					Content:   note.Text,
					MediaType: "text/x.misskeymarkdown",
				},
				Published:  note.CreatedAt,
				To:         []string{fmt.Sprintf("%s/followers", actorUrl)},
				Cc:         []string{"https://www.w3.org/ns/activitystreams#Public"},
				InReplyTo:  nil,
				Attachment: []string{},
				Sensitive:  false,
				Tag:        []string{},
			},
		}

		err := services.AddCreateActivityToDeliverQueue(queue.DeliverJob{
			UserId:         note.UserId,
			TargetInbox:    inbox,
			CreateActivity: activity,
		})

		if err != nil {
			logger.ErrorWithDetail("could not create task: %v", err)
		}
	}

	return nil
}

func Create(note *entities.Note) error {
	userService := baselib.UserService{
		LocalOnly: false,
	}

	author, err := userService.FindOne(note.UserId)
	if err != nil {
		return err
	}

	if author.Id == "" {
		return system.UserNotFound
	} else if author.IsSuspended {
		return system.UserSuspended
	}

	// サイレンス
	if author.IsSilenced && note.Visibility == "public" {
		note.Visibility = "home"
	}

	// にゃーん
	if author.IsCat {
		note.Text = utils.Nyaize(note.Text)
	}

	database, err := db.GetGormEngine()
	if err != nil {
		return err
	}
	sql := database.Table("note")

	err = sql.Create(&note).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = addToDeliverQueue(note)
	if err != nil {
		logger.ErrorWithDetail("Failed to add job", err)
		return err
	}

	return nil
}

func FindOne(id string) (*entities.Note, error) {
	var result entities.Note
	cacheKey := "note/" + id

	cacheExist := kv.GetKvCache(cacheKey, &result)

	if !cacheExist {
		engine, err := db.GetEngine()
		if err != nil {
			return &result, err
		}
		sql := engine.Table("note")
		sql.Where("id = ?", id)

		_, err = sql.Get(&result)
		if err != nil {
			return &result, err
		}

		if result.Id == "" {
			return &result, system.NoteNotFound
		}

		kv.StoreKvCache(cacheKey, result)
	}

	return &result, nil
}

func FindAllAndMap(noteIds []string) (*map[string]entities.Note, error) {
	if len(noteIds) == 0 {
		panic(system.InvalidParamsOnServiceCall)
	}

	var result map[string]entities.Note
	cacheKey := "note/" + fmt.Sprintf("%v", noteIds)

	cacheExists := kv.GetKvCache(cacheKey, &result)
	if !cacheExists {
		engine, err := db.GetEngine()
		if err != nil {
			return nil, err
		}

		var notes []entities.Note

		sql := engine.Table("note")
		sql.In("id", noteIds)

		if err := sql.Find(&notes); err != nil {
			return nil, err
		}

		result = make(map[string]entities.Note)
		for _, note := range notes {
			result[note.Id] = note
		}

		kv.StoreKvCache(cacheKey, result)
	}

	return &result, nil
}

func IsExists(id string) (bool, error) {
	var result entities.Note

	engine, err := db.GetEngine()
	if err != nil {
		return false, err
	}
	sql := engine.Table("note")
	sql.Where("id = ?", id)

	_, err = sql.Get(&result)
	if err != nil {
		return false, err
	}

	if result.Id == "" {
		return false, nil
	}

	return true, nil
}

func Delete() error {
	// ToDo
	return nil
}
