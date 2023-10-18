package note_v1

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/app/repository"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := repository.GetNoteRepository().Delete(req.GetUuid())
	if err != nil {
		log.Printf("error while reading note: %s", err.Error())
		return nil, err
	}

	return new(emptypb.Empty), nil
}
