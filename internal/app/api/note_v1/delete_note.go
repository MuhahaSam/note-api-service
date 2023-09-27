package note_v1

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/app/repository"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	"github.com/google/uuid"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*emptypb.Empty, error) {
	err := repository.GetNoteRepository().Delete(uuid.UUID(req.Uuid.Value))
	if err != nil {
		log.Printf("error while reading note: %s", err.Error())
		return nil, err
	}

	return new(emptypb.Empty), nil
}
