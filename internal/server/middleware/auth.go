package middleware

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/services/baselib"
)

type apiRequest struct {
	Token string `json:"i"`
}

func Auth(ctx *fiber.Ctx) error {
	var tokenString string

	// Misskeyのclientがcontent-typeをtext/plain;charset=UTF-8でPOSTしてくるのでfiberのBodyParserを使ったパースに失敗する
	// 応急処置として文字列として読んでそれをjsonでパースする
	req := apiRequest{}

	err := json.Unmarshal(ctx.Body(), &req)
	if err != nil {
		logger.Debug(err.Error())
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	tokenString = req.Token

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil // Verify signing key
	})

	if err != nil || !token.Valid {
		return ctx.Status(403).JSON(fiber.Map{
			"error": "Invalid token. Please re-login.",
		})
	}

	// 色々チェック
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)
	accountToken := claims["accountToken"].(string)

	if userId == "" || accountToken == "" {
		panic("JWT token was successfully decrypted but claims are invalid")
	}

	userService := baselib.UserService{
		LocalOnly: true,
	}

	i, err := userService.FindOne(userId)
	if err != nil {
		panic(err)
	}

	// 失効したトークンなら弾く
	if i.Token != accountToken {
		return ctx.Status(403).JSON(fiber.Map{
			"error": "Invalid account token. Please re-login.",
		})
	}

	if i.IsSuspended {
		return ctx.Status(403).JSON(fiber.Map{
			"error": "Your account has been suspended.",
			"id":    "a8c724b3-6e9c-4b46-b1a8-bc3ed6258370",
			"code":  "YOUR_ACCOUNT_SUSPENDED",
		})
	}

	// Store the user information in locals for later use
	ctx.Locals("user", token)

	return ctx.Next()
}
