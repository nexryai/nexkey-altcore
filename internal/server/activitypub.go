package server

import (
	"encoding/json"
	"github.com/go-fed/httpsig"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"lab.sda1.net/nexryai/altcore/internal/activitypub"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/server/middleware"
	"net/http"
	"regexp"
	"strings"
)

func inbox(ctx *fiber.Ctx) error {
	log := logger.GetLogger("inbox")

	log.Info("inbox requested")
	hr := new(http.Request)

	log.Debug("converting request")
	err := fasthttpadaptor.ConvertRequest(ctx.Context(), hr, true)
	if err != nil {
		logger.ErrorWithDetail("failed to convert request", err)
		panic(err)
	}

	// https://pkg.go.dev/github.com/valyala/fasthttp/fasthttpadaptor によれば
	// "The http.Request must not be used after the fasthttp handler has returned! Memory in use by the http.Request will be reused after your handler has returned! "
	// らしいのでポインタから中身をコピーしておく（ハンドラ関数から戻ってもジョブプロセッサーで使用するため）
	var request http.Request
	request = *hr

	log.Debug("verifying signature")
	signatureVerifier, err := httpsig.NewVerifier(&request)
	if err != nil {
		log.ErrorWithDetail("failed to create signature verifier", err)
		return ctx.SendStatus(400)
	}

	if signatureVerifier.KeyId() == "" {
		return ctx.SendStatus(401)
	}

	log.Debug("verifying digest")
	digest := ctx.Get("Digest")
	if digest == "" {
		logger.Warn("Invalid digest (empty)")
		return ctx.SendStatus(401)
	}

	match := regexp.MustCompile(`^([0-9A-Za-z-]+)=(.+)$`).FindStringSubmatch(digest)
	if match == nil {
		logger.Warn("Invalid digest (match == nil)")
		return ctx.SendStatus(401)
	}

	digestAlgo := match[1]
	//digestExpected := match[2]

	if strings.ToUpper(digestAlgo) != "SHA-256" {
		// アルゴリズムをサポートしていない
		logger.Warn("digestAlgo is not supported")
		return ctx.SendStatus(401)
	}

	/*
		hash := sha256.New()
		hash.Write(ctx.Request().Body())
		digestActual := base64.StdEncoding.EncodeToString(hash.Sum(nil))

		if digestExpected != digestActual {
			// 不正なダイジェスト
			logger.Warn("Invalid digest (digestExpected != digestActual)")
			ctx.Status(401)
			return nil
		}

		if ctx.Get("host") != config.Host {
			logger.Warn("Invalid host header")
			ctx.Status(400)
			return nil
		}*/

	// アクティビティのTypeによってobjectの型が変わるのでアクティビティの種類を判別してから構造体にマップして処理する
	log.Debug("unmarshalling activity")
	var unknownActivity = make(map[string]interface{})

	err = json.Unmarshal(ctx.Body(), &unknownActivity)
	if err != nil {
		logger.Debug(err.Error())
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid JSON-LD format",
		})
	}

	// デバッグ用
	switch unknownActivity["type"] {
	case "Follow":
		var activity activitypub.FollowActivity
		err := json.Unmarshal(ctx.Body(), &activity)
		if err != nil {
			logger.ErrorWithDetail("failed to unmarshal activity", err)
			return ctx.SendStatus(400)
		}

		err = activitypub.ProcessFollowActivity(activity)
		if err != nil {
			logger.ErrorWithDetail("failed to process activity", err)
			return ctx.SendStatus(500)
		}

	case "Create":
		var activity activitypub.CreateActivity
		err := json.Unmarshal(ctx.Body(), &activity)
		if err != nil {
			logger.ErrorWithDetail("failed to unmarshal activity", err)
			return ctx.SendStatus(400)
		}

		err = activitypub.ProcessCreateActivity(activity)
		if err != nil {
			logger.ErrorWithDetail("failed to process activity", err)
			return ctx.SendStatus(500)
		}

	case "Accept":
		logger.Info("Follow accepted")
	case "Undo":
		logger.Info("Unfollowed")
	default:
		return ctx.SendStatus(400)
	}

	return ctx.SendStatus(202)
}

func MkActivityPubRouter(app *fiber.App) {
	forceSignature := app.Group("/test", middleware.VerifySignature)
	forceSignature.Get("/debug", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"status": "OK",
		})
	})

	app.Post("/inbox", inbox)
	app.Post("/users/*/inbox", inbox)

	app.Get("/users/:userId", func(ctx *fiber.Ctx) error {
		person, err := activitypub.RenderPerson(ctx.Params("userId"))
		if err != nil {
			return ctx.SendStatus(500)
		} else if person.Id == "" {
			return ctx.SendStatus(404)
		}

		return ctx.JSON(person)
	})
}
