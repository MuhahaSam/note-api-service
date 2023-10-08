package entity

import "github.com/google/uuid"

type Entity interface {
}

type NoteEntity struct {
	Entity
	Id     uuid.UUID
	Author string
	Title  string
	Text   string
}
