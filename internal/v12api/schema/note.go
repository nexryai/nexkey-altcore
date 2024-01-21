package schema

import "time"

type Note struct {
	Id           string                 `json:"id"`
	UserId       string                 `json:"userId"`
	Visibility   string                 `json:"visibility"`
	LocalOnly    bool                   `json:"localOnly"`
	Text         string                 `json:"text"`
	CreatedAt    time.Time              `json:"createdAt"`
	User         User                   `json:"user"`
	Cw           *string                `json:"cw"`
	RenoteCount  uint                   `json:"renoteCount"`
	RepliesCount uint                   `json:"repliesCount"`
	Reactions    map[string]interface{} `json:"reactions"`
	FileIds      []string               `json:"fileIds"`
	Files        []string               `json:"files"`
	Uri          string                 `json:"uri"`
	Reply        *Note                  `json:"reply"`
	ReplyId      string                 `json:"replyId"`
	Renote       *Note                  `json:"renote"`
	RenoteId     string                 `json:"renoteId"`
}
