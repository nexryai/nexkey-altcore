package server

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"html/template"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/instance"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"os"
)

type ClientManifest struct {
	InitTS struct {
		File string `json:"file"`
	} `json:"src/init.ts"`
}

func GetClientManifest() *ClientManifest {
	body, err := os.ReadFile("built/_client_dist_/manifest.json")
	if err != nil {
		logger.FatalWithDetail("Failed to read client manifest.", err)
		panic(err)
	}

	// JSONデコード
	var manifest ClientManifest
	err = json.Unmarshal(body, &manifest)
	if err != nil {
		logger.FatalWithDetail("Failed to parse client manifest.", err)
		panic(err)
	}

	return &manifest
}

func MkClientRouter(app *fiber.App, manifest ClientManifest) {
	bootJS, err := os.ReadFile("packages/backend/built/server/web/boot.js")
	if err != nil {
		panic(err)
	}

	styleCSS, err := os.ReadFile("packages/backend/built/server/web/style.css")
	if err != nil {
		panic(err)
	}

	// 静的ファイルをサーブする
	app.Static("/assets", "built/_client_dist_")
	app.Static("/assets/locales", "built/_client_dist_/locales")
	app.Static("/assets/tabler-icons", "built/_client_dist_/tabler-icons")
	app.Static("/assets/tabler-icons/fonts", "built/_client_dist_/tabler-icons/fonts")
	app.Static("/twemoji", "packages/client/node_modules/@discordapp/twemoji/dist")
	app.Static("/", "assets")

	app.Get("/proxy/*", func(ctx *fiber.Ctx) error {
		url := ctx.Query("url")
		isAvatar := ctx.Query("avatar") == "1"
		isThumbnail := ctx.Query("thumbnail") == "1"
		isEmoji := ctx.Query("emoji") == "1"
		isTicker := ctx.Query("ticker") == "1"

		if url == "" {
			return ctx.SendStatus(400)
		}

		proxyQuery := fmt.Sprintf("url=%s", url)
		if isAvatar {
			proxyQuery += "&avatar=1"
		} else if isThumbnail {
			proxyQuery += "&thumbnail=1"
		} else if isEmoji {
			proxyQuery += "&emoji=1"
		} else if isTicker {
			proxyQuery += "&ticker=1"
		}

		return ctx.Redirect(config.MediaProxy + "?" + proxyQuery)
	})

	app.Get("/debug", func(ctx *fiber.Ctx) error {
		return ctx.Render("debug", fiber.Map{
			"Hostname":   config.Host,
			"URL":        config.URL,
			"ClientPath": config.Client,
			"InitScript": manifest.InitTS.File,
		})
	})

	// 以下catch-allなので必ずルーターの最後に来るようにする

	// 存在しないパスへのPOSTは404
	app.Post("/*", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(404)
	})

	// 連合関係でもAPIでもstaticでもない場合
	app.Get("/*", func(ctx *fiber.Ctx) error {
		meta := instance.GetInstanceMeta()

		// クライアントローダーをレンダリングする
		return ctx.Render("base", fiber.Map{
			"ClientEntryScript": manifest.InitTS.File,
			"ThemeColor":        meta.ThemeColor,
			"InstanceName":      meta.Name,
			"Icon":              meta.IconUrl,
			"Desc":              meta.Description,
			"Title":             "Nexkey",
			"bootLoader":        template.JS(bootJS),
			"initCSS":           template.CSS(styleCSS),
		})
	})
}
