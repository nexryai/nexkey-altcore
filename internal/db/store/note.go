package store

import (
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
	"time"
)

type NoteStore struct{}

func (store *NoteStore) Create(note *entities.Note) error {
	return nil
}

func (store *NoteStore) Delete() {

}

func (store *NoteStore) FindOne() {

}

func (store *NoteStore) FindMany(args struct {
	UserIds   []string
	UntilDate time.Time
	Limit     uint
	LocalOnly bool
}) (*entities.Note, error) {
	return nil, nil
}
