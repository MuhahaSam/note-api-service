package note_v1

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/app/repository"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	note, err := repository.GetNoteRepository().Get(ctx, req.GetUuid())
	if err != nil {
		log.Printf("error while reading note: %s", err.Error())
		return nil, err
	}

	return &desc.GetResponse{
		Uuid:   note.Id,
		Title:  note.Title,
		Author: note.Author,
		Text:   note.Text,
	}, nil
}
