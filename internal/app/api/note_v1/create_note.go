package note_v1

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/app/repository"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	id, err := repository.GetNoteRepository().Create(ctx, req)
	if err != nil {
		log.Printf("error while creating note: %s", err.Error())
		return nil, err
	}

	return &desc.CreateNoteResponse{
		Uuid: &desc.UUID{Value: id[:]},
	}, nil
}
