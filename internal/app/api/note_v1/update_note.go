package note_v1

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/app/repository"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	"github.com/google/uuid"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*emptypb.Empty, error) {
	err := repository.GetNoteRepository().Update(uuid.UUID(req.GetUuid().Value), req.GetUpdateBody())
	if err != nil {
		log.Printf("error while update note: %s", err.Error())
		return nil, err
	}

	return new(emptypb.Empty), nil
}
