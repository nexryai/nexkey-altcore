package account

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/db"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
	instanceCore "lab.sda1.net/nexryai/altcore/internal/services/baselib"
	apiCore "lab.sda1.net/nexryai/altcore/internal/v12api/core"
	"lab.sda1.net/nexryai/altcore/internal/v12api/schema"
	"strings"
)

func GetAccountInfo(ctx *fiber.Ctx) error {
	userId := apiCore.GetUserId(ctx)

	userService := instanceCore.UserService{
		LocalOnly: true,
	}

	myUserInfo, err := userService.FindOne(userId)
	if err != nil {
		logger.ErrorWithDetail("Failed to get user info", err)
		return ctx.Status(500).SendString("Failed to get user info")
	} else if myUserInfo.Id == "" {
		return ctx.Status(404).SendString("User not found")
	}

	myProfile, err := userService.GetProfile(userId)
	if err != nil {
		logger.ErrorWithDetail("Failed to get user profile", err)
		return ctx.Status(500).SendString("Failed to get user info")
	} else if myProfile.Id == "" {
		return ctx.Status(500).SendString("Broken user profile")
	}

	if myUserInfo.Avatar == nil {
		myUserInfo.Avatar = &entities.DriveFile{}
	}

	if myUserInfo.Banner == nil {
		myUserInfo.Banner = &entities.DriveFile{}
	}

	return ctx.JSON(&schema.MyAccount{
		Id:             userId,
		Name:           myUserInfo.Name,
		Username:       myUserInfo.Username,
		Host:           nil,
		IsAdmin:        myUserInfo.IsAdmin,
		Location:       myProfile.Location,
		IsCat:          myUserInfo.IsCat,
		AvatarId:       myUserInfo.AvatarId,
		AvatarURL:      myUserInfo.Avatar.URL,
		AvatarBlurhash: myUserInfo.Avatar.BlurHash,
		BannerId:       myUserInfo.BannerId,
		BannerUrl:      myUserInfo.Banner.URL,
		BannerBlurhash: myUserInfo.Banner.BlurHash,
	})
}

func GetRegistry(ctx *fiber.Ctx) error {
	// ToDo: 全体的にDBが意味不明な設計なのでリファクタリングする
	userId := apiCore.GetUserId(ctx)
	var items []entities.UserRegistryItem

	engine, err := db.GetEngine()
	if err != nil {
		panic(err)
	}

	sql := engine.Table("registry_item")
	sql.Where("\"userId\" = ?", userId)

	err = sql.Find(&items)
	if err != nil {
		logger.FatalWithDetail("Failed to get registry", err)
		panic(err)
	}

	// Nexkey Super Flexible Json Encode Technology
	// 型がめちゃくちゃでfiberにjsonとして扱わせるとpanicになるので気合でjsonにする
	// それくらいDBの設計がめちゃくちゃなので早くマイグレーション書いて何とかする
	ctx.Set("Content-Type", "application/json")
	var jsonStr = "{"

	for _, x := range items {
		jsonStr += fmt.Sprintf("\"%s\": %s, ", x.Key, x.Value)
	}

	jsonStr += "}"

	// 最後のコンマを消す
	lastCommaIndex := strings.LastIndex(jsonStr, ",")
	if lastCommaIndex != -1 {
		result := jsonStr[:lastCommaIndex] + jsonStr[lastCommaIndex+1:]
		return ctx.Status(200).Send([]byte(result))
	}

	return ctx.Status(200).Send([]byte(jsonStr))
}
