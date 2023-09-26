package entity

import "github.com/google/uuid"

type NoteEntity struct {
	Id     uuid.UUID
	Author string
	Title  string
	Text   string
}