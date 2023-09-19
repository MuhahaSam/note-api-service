package note_v1

import desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"

type Note struct {
	desc.NoteServiceServer
}

func NewNote() *Note {
	return &Note{}
}
