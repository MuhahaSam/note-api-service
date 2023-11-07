package note_v1

import (
	"github.com/MuhahaSam/golangPractice/internal/service/note"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteServiceServer

	noteService *note.Service
}

func NewNoteV1(eventService *note.Service) *Implementation {
	return &Implementation{
		desc.UnimplementedNoteServiceServer{},

		eventService,
	}
}
