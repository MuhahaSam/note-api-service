package note_v1

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/app/repository"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	"github.com/google/uuid"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	note, err := repository.GetNoteRepository().Read(ctx, uuid.MustParse(req.Uuid))
	if err != nil {
		log.Printf("error while reading note: %s", err.Error())
		return nil, err
	}

	return &desc.GetNoteResponse{
		Uuid:   note.Id.String(),
		Title:  note.Title,
		Author: note.Author,
		Text:   note.Text,
	}, nil
}
