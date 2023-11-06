package note_v1

import (
	"context"
	"log"

	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (n *Implementation) DeleteNote(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := n.noteService.Delete(ctx, req)
	if err != nil {
		log.Printf("error while reading note: %s", err.Error())
		return nil, err
	}

	return new(emptypb.Empty), nil
}
