package note

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"lab.sda1.net/nexryai/altcore/internal/core"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
	noteService "lab.sda1.net/nexryai/altcore/internal/services/xnote"
	apiCore "lab.sda1.net/nexryai/altcore/internal/v12api/core"
	"lab.sda1.net/nexryai/altcore/internal/v12api/schema"
	"time"
)

type createNoteResp struct {
	CreatedNote *schema.Note `json:"createdNote"`
}

func CreateNote(ctx *fiber.Ctx) error {
	req := schema.Note{
		UserId: apiCore.GetUserId(ctx),
	}

	apiCore.ParseRequest(ctx, &req)

	noteId := core.GenId()
	note := &entities.Note{
		Id:         noteId,
		UserId:     req.UserId,
		Visibility: req.Visibility,
		Text:       req.Text,
		CreatedAt:  time.Now(),
		LocalOnly:  true,
		//Reactions:  []uint8{},
		UserHost: "",
		Uri:      fmt.Sprintf("%s/notes/%s", config.URL, noteId),
	}

	// リプライ
	if req.ReplyId != "" {
		exists, err := noteService.IsExists(req.ReplyId)
		if err != nil {
			return ctx.SendStatus(500)
		} else if !exists {
			// リプライ先のノートがない
			return ctx.SendStatus(400)
		}

		note.ReplyId = req.ReplyId
	}

	err := noteService.Create(note)
	if err != nil {
		panic(err)
	}

	return ctx.JSON(createNoteResp{
		CreatedNote: &req,
	})
}
