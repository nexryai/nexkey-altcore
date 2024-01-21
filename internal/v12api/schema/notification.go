package schema

import "time"

type Notification struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Type      string    `json:"type"`
	IsRead    bool      `json:"isRead"`
	Note      Note      `json:"note"`
	User      User      `json:"user"`
	Reaction  string    `json:"reaction"`
}
