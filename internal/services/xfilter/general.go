package xfilter

import (
	"lab.sda1.net/nexryai/altcore/internal/core/enum"
	"lab.sda1.net/nexryai/altcore/internal/core/utils"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
)

type Cache struct {
	// 任意: なかったらDBに取りに行く（既に上位の関数で読み込まれてる場合に渡せばその情報を再利用してクエリを減らせる）
	CachedFollowees *[]string
	CachedMutees    *[]string
	CachedBlockers  *[]string
}

func checkVisibility(noteVisibility enum.NoteVisibility, userId string, cache *Cache) bool {
	switch noteVisibility {
	case enum.NoteVisibilityPublic:
		return true
	case enum.NoteVisibilityHome:
		return true
	case enum.NoteVisibilityFollowersOnly:
		return utils.ContainsString(*cache.CachedFollowees, userId)
	case enum.NoteVisibilityDirect:
		return false
	default:
		return false
	}
}

func isMuted(note *entities.Note, userId string, cache *Cache) bool {
	return utils.ContainsString(*cache.CachedMutees, note.UserId)
}

func isBlocked(note *entities.Note, userId string, cache *Cache) bool {
	return utils.ContainsString(*cache.CachedBlockers, note.UserId)
}

func isOk(note *entities.Note, userId string, cache *Cache) bool {
	return checkVisibility(enum.NoteVisibility(note.Visibility), userId, cache) && !isMuted(note, userId, cache) && !isBlocked(note, userId, cache)
}

func FilterNotes(notes *[]entities.Note, userId string, cache *Cache) *[]entities.Note {
	for _, n := range *notes {
		if isOk(&n, userId, cache) {
			if n.Reply != nil {
				if !isOk(n.Reply, userId, cache) {
					continue
				}
			}

			if n.Renote != nil {
				if !isOk(n.Renote, userId, cache) {
					continue
				}
			}

			*notes = append(*notes, n)
		}
	}

	return notes
}
