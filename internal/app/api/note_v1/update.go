package note_v1

import (
	"context"
	"log"

	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (n *Implementation) UpdateNote(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := n.noteService.Update(ctx, req)
	if err != nil {
		log.Printf("error while update note: %s", err.Error())
		return nil, err
	}

	return new(emptypb.Empty), nil
}
