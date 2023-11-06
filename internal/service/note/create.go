package note

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/converter"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

func (n *Service) Create(ctx context.Context, req *desc.CreateRequest) (*string, error) {
	noteEntity := converter.ToNoteModel(req)
	id, err := n.noteRepository.Create(ctx, noteEntity)
	if err != nil {
		log.Printf("error while creating note: %s", err.Error())
		return nil, err
	}
	var uuid string = id.String()

	return &uuid, nil
}
