package note_v1

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/app/db"
	"github.com/MuhahaSam/golangPractice/internal/app/repository"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	db.GetDbModuleInstance().Connect()
	note, err := repository.GetNoteRepository().Read(req.Index)
	if err != nil {
		log.Fatalf("error while reading note: %s", err.Error())
	}

	defer db.GetDbModuleInstance().Close()

	return &desc.GetNoteResponse{
		Index:  note.Index,
		Title:  note.Title,
		Author: note.Author,
		Text:   note.Text,
	}, nil

}
