package boot

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/server"
	"strings"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	// APIへのアクセスならJSONでエラーを返す
	if strings.HasPrefix(ctx.Path(), "/api") {
		return ctx.Status(code).JSON(ErrorResponse{Error: "Internal Sever Error"})
	} else {
		return ctx.Status(code).Render("error", fiber.Map{
			"isDebugMode": true,
		})
	}
}

func StartWebServer() {
	logger.ProgressInfo("Configuring router...")
	engine := html.New("./internal/server/client-loader/views", ".gohtml")

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Views:                 engine,
		ErrorHandler:          errorHandler,
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	server.MkV12ApiRouter(app)
	server.MkWellKnownRouter(app)

	server.MkActivityPubRouter(app)

	// 連合関係でもAPIでもないcatch-allなハンドラーなので最後に設置
	manifest := server.GetClientManifest()
	server.MkClientRouter(app, *manifest)

	logger.ProgressOk()
	fmt.Print("\n")

	listenOn := fmt.Sprintf("%s:%d", "0.0.0.0", config.Port)

	logger.Info(fmt.Sprintf("Start listening on %s", listenOn))
	err := app.Listen(listenOn)
	if err != nil {
		logger.FatalWithDetail("Failed to listen", err)
	}
}
