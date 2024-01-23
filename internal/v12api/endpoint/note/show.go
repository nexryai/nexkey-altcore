package note

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/core/system"
	noteService "lab.sda1.net/nexryai/altcore/internal/services/xnote"
	apiCore "lab.sda1.net/nexryai/altcore/internal/v12api/core"
	"lab.sda1.net/nexryai/altcore/internal/v12api/schema"
)

type showNoteParam struct {
	UserId string `json:"userId"`
	NoteId string `json:"noteId"`
}

func ShowNote(ctx *fiber.Ctx) error {
	req := showNoteParam{
		UserId: apiCore.GetUserId(ctx),
	}

	apiCore.ParseRequest(ctx, &req)

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
