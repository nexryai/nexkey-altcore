package server

import (
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/server/middleware"
	"lab.sda1.net/nexryai/altcore/internal/v12api"
	streamingApi "lab.sda1.net/nexryai/altcore/internal/v12api/streaming"
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
	v12api.RoutePrivateApi(protectedApi)
}

func MkV12ApiRouter(app *fiber.App) {
	// ストリーミング
	app.Get("/streaming", websocket.New(func(ctx *websocket.Conn) {
		streamingApi.HandleStreamingApi(ctx)
	}))

	v12api.RoutePublicApi(app)
	mkPrivateApi(app)
}
