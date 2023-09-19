package note_v1

import (
	db "github.com/MuhahaSam/golangPractice/internal/app/db"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

type Note struct {
	desc.NoteServiceServer
}

func NewNote() *Note {

	db.GetDbModuleInstance().Connect()

	return &Note{}
}
