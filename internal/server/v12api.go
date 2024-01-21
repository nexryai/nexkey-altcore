package server

import (
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/server/middleware"
	"lab.sda1.net/nexryai/altcore/internal/v12api"
	"lab.sda1.net/nexryai/altcore/internal/v12api/schema"
	"math/rand"
)

func getUserId(ctx *fiber.Ctx) string {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	// 復号化できてるはずなのにuserIdがない場合panic
	if userId == "" {
		panic("JWT token was successfully decrypted but claims are invalid")
	}

	return userId
}

func parseRequest(ctx *fiber.Ctx, param interface{}) interface{} {
	err := json.Unmarshal(ctx.Body(), &param)
	if err != nil {
		logger.Debug(err.Error())
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	return nil
}

func mkPrivateApi(app *fiber.App) {
	protectedApi := app.Group("/api", middleware.Auth)

	// /i
	protectedApi.Post("/i", func(ctx *fiber.Ctx) error {
		return v12api.GetAccountInfo(ctx)
	})
	protectedApi.Post("/i/registry/get-all", func(ctx *fiber.Ctx) error {
		return v12api.GetRegistry(ctx)
	})

	// /notes
	protectedApi.Post("/notes/create", func(ctx *fiber.Ctx) error {
		return v12api.CreateNote(ctx)
	})
	protectedApi.Post("/notes/timeline", func(ctx *fiber.Ctx) error {
		return v12api.GetHomeTimeline(ctx)
	})

	// 寂しいのでダミーでとりあえず
	protectedApi.Post("/get-online-users-count", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"count": rand.Intn(9999),
			"dummy": true,
		})
	})
}

func MkV12ApiRouter(app *fiber.App) {
	// ストリーミング
	app.Get("/streaming", websocket.New(func(ctx *websocket.Conn) {
		v12api.HandleStreamingApi(ctx)
	}))

	// 認証が要らないエンドポイント
	app.Post("/api/meta", func(ctx *fiber.Ctx) error {
		return ctx.JSON(v12api.GetMeta())
	})

	app.Post("/api/signin", func(ctx *fiber.Ctx) error {
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

		resp, err := v12api.SignIn(req)

		switch err {
		case v12api.ErrorPasswordInvalid:
			return ctx.Status(403).JSON(fiber.Map{
				"error": "Invalid password",
				"id":    "932c904e-9460-45b7-9ce6-7ed33be7eb2c",
			})
		case v12api.ErrorUserNotFound:
			return ctx.Status(400).JSON(fiber.Map{
				"error": "Invalid username",
				"id":    "6cc579cc-885d-43d8-95c2-b8c7fc963280",
			})
		case nil:
			// エラーがないなら成功
			return ctx.JSON(resp)
		default:
			// どれにも当てはまらないならとりあえず500
			return ctx.SendStatus(500)
		}

	})

	mkPrivateApi(app)
}
