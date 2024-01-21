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

func GetHomeTimeline(ctx *fiber.Ctx) error {
	req := getHomeTimelineParam{
		UserId: getUserId(ctx),
	}

	parseRequest(ctx, &req)

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

	database, _ := db.GetGormEngine()

	err = database.Preload("User").
		Preload("User.Avatar").
		Preload("User.Banner").
		Preload("Reply").
		Preload("Reply.User").
		Preload("Reply.User.Avatar").
		Preload("Reply.User.Banner").
		Where("note.\"userId\" IN (?) AND note.visibility IN (?)", userIds, visibilities).
		Limit(10).
		Order("note.\"id\" DESC").
		Find(&notes).
		Error

	if err != nil {
		logger.ErrorWithDetail("Failed to get notes", err)
		return ctx.SendStatus(500)
	}

	if len(*notes) == 0 {
		return ctx.JSON(&notes)
	}

	var resp []schema.Note

	for _, n := range *notes {
		u := n.User

		if u.Avatar == nil {
			u.Avatar = &entities.DriveFile{}
		}

		avatarUrl := u.Avatar.URL
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
			Cw:        n.Cw,
		}

		if n.Reply != nil {
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

		resp = append(resp, newNote)
	}

	return ctx.JSON(&resp)
}
