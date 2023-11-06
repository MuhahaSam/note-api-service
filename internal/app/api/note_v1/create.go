package note_v1

import (
	"context"
	"log"

	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

func (n *Implementation) CreateNote(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	id, err := n.noteService.Create(ctx, req)

	if err != nil {
		log.Printf("error while reading note: %s", err.Error())
		return nil, err
	}

	return &desc.CreateResponse{
		Uuid: *id,
	}, nil
}
