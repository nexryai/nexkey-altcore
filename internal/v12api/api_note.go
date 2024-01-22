package v12api

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"lab.sda1.net/nexryai/altcore/internal/core"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/core/system"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
	noteService "lab.sda1.net/nexryai/altcore/internal/services/xnote"
	"lab.sda1.net/nexryai/altcore/internal/v12api/schema"
	"time"
)

type showNoteParam struct {
	UserId string `json:"userId"`
	NoteId string `json:"noteId"`
}

type createNoteResp struct {
	CreatedNote *schema.Note `json:"createdNote"`
}

func ShowNote(ctx *fiber.Ctx) error {
	req := showNoteParam{
		UserId: getUserId(ctx),
	}

	parseRequest(ctx, &req)

	note, err := noteService.FindOne(req.NoteId)
	if errors.Is(err, system.NoteNotFound) {
		return ctx.SendStatus(404)
	} else if err != nil {
		logger.ErrorWithDetail("failed to find note", err)
		return ctx.SendStatus(500)
	}

	resp := schema.Note{
		Id:         note.Id,
		UserId:     note.UserId,
		Visibility: note.Visibility,
		Text:       note.Text,
		CreatedAt:  note.CreatedAt,
		LocalOnly:  note.LocalOnly,
		//Reactions:  note.Reactions,
		Uri: note.Uri,
		Cw:  note.Cw,
	}

	return ctx.JSON(resp)
}

func CreateNote(ctx *fiber.Ctx) error {
	req := schema.Note{
		UserId: getUserId(ctx),
	}

	parseRequest(ctx, &req)

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