package note

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/model"
)

func (n *Service) Get(ctx context.Context, uuid string) (*model.NoteEntity, error) {
	note, err := n.noteRepository.Get(ctx, uuid)
	if err != nil {
		log.Printf("error while reading note: %s", err.Error())
		return nil, err
	}

	return note, nil
}
