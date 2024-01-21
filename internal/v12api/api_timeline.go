package v12api

import (
	"github.com/gofiber/fiber/v2"
	"lab.sda1.net/nexryai/altcore/internal/core/enum"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/db"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
	"lab.sda1.net/nexryai/altcore/internal/services/baselib"
	"lab.sda1.net/nexryai/altcore/internal/v12api/schema"
	"log"
)

type getHomeTimelineParam struct {
	// "i" パラメータから自動生成する
	UserId  string
	Limit   uint   `json:"limit"`
	UntilId string `json:"untilId"`
}

func findNoteById(notes []entities.Note, targetId string) *entities.Note {
	for _, note := range notes {
		if note.Id == targetId {
			return &note
		}
	}
	return nil
}

func GetHomeTimeline(ctx *fiber.Ctx) error {
	req := getHomeTimelineParam{
		UserId: getUserId(ctx),
	}

	parseRequest(ctx, &req)

	// xormエンジンの初期化
	//engine, err := db.GetEngine()

	followService := baselib.FollowService{
		Type:   enum.Followees,
		UserId: req.UserId,
	}

	followees, err := followService.FindAll()
	if err != nil {
		log.Fatal(err)
	}

	userIds := followees
	// 自分のノート
	userIds = append(userIds, req.UserId)
	visibilities := []string{"public", "home"}

	var notes *[]entities.Note

	db, _ := db.GetGormEngine()

	err = db.Preload("User").
		Preload("Avatar").
		Preload("Banner").
		Preload("ReplyNotes").
		Preload("ReplyNotes.User").
		Preload("ReplyNotes.Avatar").
		Preload("ReplyNotes.Banner").
		Where("note.userId IN (?) AND note.visibility IN (?)", userIds, visibilities).
		Limit(10).
		Find(&notes).
		Error

	if err != nil {
		logger.ErrorWithDetail("Failed to get notes", err)
		return ctx.SendStatus(500)
	}

	/*
		sql := engine.Table("note")
		sql.Select("note.*, \"user\".*, \"avatar\".*").
			Join("INNER", "user", "\"user\".\"id\" = note.\"userId\"").
			Join("LEFT", "\"drive_file\" as \"avatar\"", "\"avatar\".\"id\" = \"user\".\"avatarId\"").
			Join("LEFT", "\"drive_file\" as \"banner\"", "\"banner\".\"id\" = \"user\".\"bannerId\"").
			// FIXME: なんかうごかん
			Join("LEFT", "note as reply", "\"note\".\"id\" = reply.\"replyId\" AND reply.\"replyId\" IS NOT NULL").
			Join("INNER", "user as \"user_reply\"", "\"user_reply\".\"id\" = reply.\"userId\"").
			Join("LEFT", "\"drive_file\" as \"avatar_reply\"", "\"avatar_reply\".\"id\" = \"user_reply\".\"avatarId\"").
			Join("LEFT", "\"drive_file\" as \"banner_reply\"", "\"banner_reply\".\"id\" = \"user_reply\".\"bannerId\"").
			In("note.userId", userIds).
			In("note.visibility", visibilities)

		if req.UntilId != "" {
			sql.Where("note.id <= ?", req.UntilId)
		}

		sql.Limit(int(req.Limit), 0)
		sql.OrderBy("note.\"createdAt\" DESC")

		logger.Debug(fmt.Sprintf("SQL: %v", sql))
		if err := sql.Find(&notes); err != nil {
			log.Fatal(err)
		}*/

	if len(*notes) == 0 {
		return ctx.JSON(&notes)
	}

	// リプライ・リノート用
	// additionalNotes = リプライ・リノート
	/*
		var additionalNotes *map[string]entities.Note
		var additionalNotesIds []string
		var additionalUsers *map[string]entities.User
		var additionalUsersIds []string
		var additionalAvatars *map[string]entities.DriveFile
		var additionalAvatarsIds []string

		for _, note := range notes {
			if note.Note.ReplyId != nil {
				additionalNotesIds = append(additionalNotesIds, *note.Note.ReplyId)
			}
		}

		if len(additionalNotesIds) > 0 {
			noteService := xnote.NoteService{}
			additionalNotes, err = noteService.FindAllAndMap(additionalNotesIds)
			if err != nil {
				logger.ErrorWithDetail("Failed to get additional notes", err)
				return ctx.SendStatus(500)
			}

			// additionalNotesのユーザーを取得
			for _, note := range *additionalNotes {
				additionalUsersIds = append(additionalUsersIds, note.UserId)
			}

			userService := baselib.UserService{}
			additionalUsers, err = userService.FindAllAndMap(additionalUsersIds)
			if err != nil {
				logger.ErrorWithDetail("Failed to get additional users", err)
				return ctx.SendStatus(500)
			}

			// additionalNotesのアバターを取得
			for _, user := range *additionalUsers {
				if user.AvatarId != "" {
					additionalAvatarsIds = append(additionalAvatarsIds, user.AvatarId)
				}
			}

			driveService := xdrive.DriveService{}
			additionalAvatars, err = driveService.FindAllAndMap(additionalAvatarsIds)
			if err != nil {
				logger.ErrorWithDetail("Failed to get additional avatars", err)
				return ctx.SendStatus(500)
			}
		}*/

	var resp []schema.Note

	for _, n := range *notes {
		u := n.User

		avatarUrl := n.User.Avatar.URL
		if n.User.Host == "" {
			avatarUrl = n.User.Avatar.Thumbnail
		} else if avatarUrl == "" {
			avatarUrl = "https://raw.githubusercontent.com/misskey-dev/misskey/develop/packages/backend/assets/user-unknown.png"
		}

		newNote := schema.Note{
			Id:         n.Id,
			Text:       n.Text,
			UserId:     n.UserId,
			Visibility: n.Visibility,
			LocalOnly:  n.LocalOnly,
			User: schema.User{
				Id:             u.Id,
				Name:           u.Name,
				Username:       u.Username,
				Host:           u.Host,
				AvatarUrl:      avatarUrl,
				AvatarBlurhash: u.Avatar.BlurHash,
				IsBot:          u.IsBot,
				IsCat:          u.IsCat,
			},
			CreatedAt: n.CreatedAt,
			Reactions: n.Reactions,
			FileIds:   []string{},
			Files:     []string{},
			Cw:        nil,
		}

		if n.ReplyId != nil {
			replyTarget := n.Reply
			replyTargetUser := n.Reply.User
			replyTargetUserAvatar := n.Reply.User.Avatar

			newNote.Reply = &schema.Note{
				Id:         replyTarget.Id,
				UserId:     replyTarget.UserId,
				Visibility: replyTarget.Visibility,
				Text:       replyTarget.Text,
				CreatedAt:  replyTarget.CreatedAt,
				LocalOnly:  replyTarget.LocalOnly,
				Reactions:  replyTarget.Reactions,
				Uri:        replyTarget.Uri,
				Cw:         replyTarget.Cw,
				User: schema.User{
					Id:             replyTargetUser.Id,
					Name:           replyTargetUser.Name,
					Username:       replyTargetUser.Username,
					Host:           replyTargetUser.Host,
					AvatarUrl:      replyTargetUserAvatar.URL,
					AvatarBlurhash: replyTargetUserAvatar.BlurHash,
					IsBot:          replyTargetUser.IsBot,
					IsCat:          replyTargetUser.IsCat,
				},
				RepliesCount: 1,
				Files:        []string{},
				FileIds:      []string{},
			}

		}

		/*
			if len(n.ReplyNotes) > 0 {
				logger.Debug("hadReply")
				newNote.ReplyId = *n.Note.ReplyId
				replyTarget := n.ReplyNotes[0].Note

				newNote.Reply = &schema.Note{
					Id:         replyTarget.Id,
					UserId:     replyTarget.UserId,
					Visibility: replyTarget.Visibility,
					Text:       replyTarget.Text,
					CreatedAt:  replyTarget.CreatedAt,
					LocalOnly:  replyTarget.LocalOnly,
					Reactions:  replyTarget.Reactions,
					Uri:        replyTarget.Uri,
					Cw:         replyTarget.Cw,
					User: schema.User{
						Id:             n.ReplyNotes[0].User.Id,
						Name:           n.ReplyNotes[0].User.Name,
						Username:       n.ReplyNotes[0].User.Username,
						Host:           n.ReplyNotes[0].User.Host,
						AvatarUrl:      n.ReplyNotes[0].Avatar.URL,
						AvatarBlurhash: n.ReplyNotes[0].Avatar.BlurHash,
						IsBot:          n.ReplyNotes[0].User.IsBot,
						IsCat:          n.ReplyNotes[0].User.IsCat,
					},
					RepliesCount: 1,
					Files:        []string{},
					FileIds:      []string{},
				}

			}*/

		resp = append(resp, newNote)
	}

	return ctx.JSON(&resp)
}
