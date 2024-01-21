package entities

import "time"

type Notification struct {
	Id              string    `xorm:"id"`
	CreatedAt       time.Time `xorm:"createdAt"`
	NotifieeId      string    `xorm:"notifieeId"`
	NotiferId       string    `xorm:"notiferId"`
	IsRead          bool      `xorm:"isRead"`
	NoteId          string    `xorm:"noteId"`
	Reaction        string    `xorm:"reaction"`
	Choice          string    `xorm:"choice"`
	FollowRequestId string    `xorm:"followRequestId"`
	Type            string    `xorm:"type"`
}
