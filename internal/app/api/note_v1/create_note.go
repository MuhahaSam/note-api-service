package note_v1

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/app/repository"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	id, err := repository.GetNoteRepository().Create(req)
	if err != nil {log.Fatalf("error while creating note: %s", err.Error())}

	return &desc.CreateNoteResponse{
		Uuid: &desc.UUID{Value: id[:]},
	}, nil
}
