package note

import (
	"context"
	"log"

	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

func (n *Service) Delete(ctx context.Context, req *desc.DeleteRequest) error {
	err := n.noteRepository.Delete(ctx, req.GetUuid())
	if err != nil {
		log.Printf("error while reading note: %s", err.Error())
		return err
	}

	return nil
}
