package note

import (
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
	"time"
)

func Create(note *entities.Note) error {
	return nil
}

func Delete() {

}

func FindOne() {

}

func FindMany(args struct {
	UserIds   []string
	UntilDate time.Time
	Limit     uint
	LocalOnly bool
}) (*entities.Note, error) {
	return nil, nil
}
