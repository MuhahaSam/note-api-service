package note

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/model"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

func (n *Service) Get(ctx context.Context, req *desc.GetRequest) (*model.NoteEntity, error) {
	note, err := n.noteRepository.Get(ctx, req.GetUuid())
	if err != nil {
		log.Printf("error while reading note: %s", err.Error())
		return nil, err
	}

	return note, nil
}
