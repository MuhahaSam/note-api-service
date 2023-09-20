package note_v1

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/app/db"
	"github.com/MuhahaSam/golangPractice/internal/app/repository"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.DeleteNoteResponse, error) {
	db.GetDbModuleInstance().Connect()
	err := repository.GetNoteRepository().Delete(req.Index)
	if err != nil {
		log.Fatalf("error while reading note: %s", err.Error())
	}

	defer db.GetDbModuleInstance().Close()

	return &desc.DeleteNoteResponse{}, nil

}
