package entities

import "time"

type Note struct {
	Id           string    `gorm:"primaryKey"`
	UserId       string    `gorm:"column:userId"`
	User         *User     `gorm:"foreignkey:Id;references:UserId"`
	Visibility   string    `gorm:"column:visibility"`
	Text         string    `gorm:"column:text"`
	CreatedAt    time.Time `gorm:"column:createdAt"`
	ReplyId      string    `gorm:"column:replyId"`
	Reply        *Note     `gorm:"foreignkey:Id;references:ReplyId"`
	RenoteId     *string   `gorm:"column:renoteId"`
	Renote       *Note     `gorm:"foreignkey:Id;references:RenoteId"`
	Cw           string    `gorm:"column:cw"`
	LocalOnly    bool      `gorm:"column:localOnly"`
	RepliesCount int       `gorm:"column:repliesCount"`
	RenoteCount  int       `gorm:"column:renoteCount"`
	Reactions    []uint8   `gorm:"column:reactions"`
	Uri          string    `gorm:"column:uri;unique"`
	Score        int       `gorm:"column:score"`
	UserHost     string    `gorm:"column:userHost"`
}

func (Note) TableName() string {
	return "note"
}
