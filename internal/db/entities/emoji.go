package entities

import (
	"github.com/lib/pq"
	"time"
)

type Emoji struct {
	Id        string         `xorm:"id"`
	UpdatedAt time.Time      `xorm:"updatedAt"`
	Name      string         `xorm:"name"`
	Host      *string        `xorm:"host"`
	OriginUrl string         `xorm:"originUrl"`
	Uri       string         `xorm:"uri"`
	Aliases   pq.StringArray `xorm:"aliases"`
	PublicUrl string         `xorm:"publicUrl"`
}
