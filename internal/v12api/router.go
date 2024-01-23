package v12api

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"lab.sda1.net/nexryai/altcore/internal/core/instance"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	accountApi "lab.sda1.net/nexryai/altcore/internal/v12api/endpoint/account"
	adminApi "lab.sda1.net/nexryai/altcore/internal/v12api/endpoint/admin"
	metaApi "lab.sda1.net/nexryai/altcore/internal/v12api/endpoint/meta"
	noteApi "lab.sda1.net/nexryai/altcore/internal/v12api/endpoint/note"
	"lab.sda1.net/nexryai/altcore/internal/v12api/schema"
	"math/rand"
)

func RoutePublicApi(publicRouter fiber.Router) {
	publicRouter.Post("/api/signin", func(ctx *fiber.Ctx) error {
		req := new(schema.SignInRequest)
		if err := ctx.BodyParser(req); err != nil {
			// Misskeyのclientがcontent-typeをtext/plain;charset=UTF-8でPOSTしてくるのでfiberのBodyParserを使ったパースに失敗する
			// 応急処置として文字列として読んでそれをjsonでパースする
			err := json.Unmarshal(ctx.Body(), req)
			if err != nil {
				logger.Debug(err.Error())
				return ctx.Status(400).JSON(fiber.Map{
					"error": "Invalid JSON format",
				})
			}
		}

		resp, err := SignIn(req)

		switch {
		case errors.Is(err, ErrorPasswordInvalid):
			return ReturnApiError(ctx, ApiErrPasswordInvalid)
		case errors.Is(err, ErrorUserNotFound):
			return ReturnApiError(ctx, ApiErrUserNotFound)
		case err == nil:
			// エラーがないなら成功
			return ctx.JSON(resp)
		default:
			// どれにも当てはまらないならとりあえず500
			return ctx.SendStatus(500)
		}

	})

	// 認証が要らないエンドポイント
	publicRouter.Post("/api/meta", func(ctx *fiber.Ctx) error {
		return ctx.JSON(metaApi.GetMeta())
	})

	// 初回セットアップ用
	if instance.ShouldCreateAdminAccount() {
		publicRouter.Post("/api/admin/accounts/create", func(ctx *fiber.Ctx) error {
			req := new(schema.SignInRequest)
			err := json.Unmarshal(ctx.Body(), req)
			if err != nil {
				logger.Debug(err.Error())
				return ctx.Status(400).JSON(fiber.Map{
					"error": "Invalid JSON format",
				})
			}

			return adminApi.CreateAdminAccount(ctx)
		})
	}
}

func RoutePrivateApi(privateRouter fiber.Router) {
	// /i
	privateRouter.Post("/i", func(ctx *fiber.Ctx) error {
		return accountApi.GetAccountInfo(ctx)
	})
	privateRouter.Post("/i/registry/get-all", func(ctx *fiber.Ctx) error {
		return accountApi.GetRegistry(ctx)
	})

	// /notes
	privateRouter.Post("/notes/create", func(ctx *fiber.Ctx) error {
		return noteApi.CreateNote(ctx)
	})
	privateRouter.Post("/notes/timeline", func(ctx *fiber.Ctx) error {
		return noteApi.GetHomeTimeline(ctx)
	})

	// 寂しいのでダミーでとりあえず
	privateRouter.Post("/get-online-users-count", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"count": rand.Intn(9999),
			"dummy": true,
		})
	})
}
