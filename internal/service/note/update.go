package note

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/converter"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

func (n *Service) Update(ctx context.Context, req *desc.UpdateRequest) error {
	noteUpdateBody := converter.ToNoteUpdateBody(req.GetUpdateBody())
	err := n.noteRepository.Update(ctx, req.GetUuid(), noteUpdateBody)
	if err != nil {
		log.Printf("error while update note: %s", err.Error())
		return err
	}

	return nil
}
