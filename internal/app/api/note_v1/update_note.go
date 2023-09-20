package note_v1

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/app/db"
	"github.com/MuhahaSam/golangPractice/internal/app/repository"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	db.GetDbModuleInstance().Connect()
	err := repository.GetNoteRepository().Update(req.GetIndex(), req.GetUpdateBody())
	if err != nil {
		log.Fatalf("error while update note: %s", err.Error())
	}

	defer db.GetDbModuleInstance().Close()

	return &desc.UpdateNoteResponse{}, nil

}
