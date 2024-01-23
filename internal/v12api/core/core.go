package core

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
)

var (
	InvalidParam = errors.New("Invalid param")
)

func GetUserId(ctx *fiber.Ctx) string {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	// 復号化できてるはずなのにuserIdがない場合panic
	if userId == "" {
		panic("JWT token was successfully decrypted but claims are invalid")
	}

	return userId
}

func ParseRequest(ctx *fiber.Ctx, param interface{}) interface{} {
	err := json.Unmarshal(ctx.Body(), &param)
	if err != nil {
		logger.Debug(err.Error())
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	return nil
}
