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
	"lab.sda1.net/nexryai/altcore/internal/services/xnote"
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

	noteService := xnote.NoteService{
		RequesterUserIdForVisibilityCheck: req.UserId,
	}

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
		Reactions:  note.Reactions,
		Uri:        note.Uri,
		Cw:         note.Cw,
	}

	if note.ReplyId != nil {
		resp.ReplyId = *note.ReplyId

		replayNote, err := noteService.FindOne(*note.ReplyId)
		if errors.Is(err, system.NoteNotFound) {
			// 閲覧権限がないノートへのリプライも表示しない
			return ctx.SendStatus(404)
		} else if err != nil {
			logger.ErrorWithDetail("failed to find note", err)
			return ctx.SendStatus(500)
		}

		resp.Reply = &schema.Note{
			Id:         replayNote.Id,
			UserId:     replayNote.UserId,
			Visibility: replayNote.Visibility,
			Text:       replayNote.Text,
			CreatedAt:  replayNote.CreatedAt,
			LocalOnly:  replayNote.LocalOnly,
			Reactions:  replayNote.Reactions,
			Uri:        replayNote.Uri,
			Cw:         note.Cw,
		}
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
		Reactions:  make(map[string]interface{}),
		UserHost:   nil,
		Uri:        fmt.Sprintf("%s/notes/%s", config.URL, noteId),
	}

	// リプライ
	if req.ReplyId != "" {
		noteService := xnote.NoteService{}
		exists, err := noteService.IsExists(req.ReplyId)
		if err != nil {
			return ctx.SendStatus(500)
		} else if !exists {
			// リプライ先のノートがない
			return ctx.SendStatus(400)
		}

		note.ReplyId = &req.ReplyId
	}

	noteService := xnote.NoteService{
		UserId: req.UserId,
	}
	err := noteService.Create(note)
	if err != nil {
		panic(err)
	}

	return ctx.JSON(createNoteResp{
		CreatedNote: &req,
	})
}
