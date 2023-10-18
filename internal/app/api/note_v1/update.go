package note_v1

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/app/repository"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := repository.GetNoteRepository().Update(req.GetUuid(), req.GetUpdateBody())
	if err != nil {
		log.Printf("error while update note: %s", err.Error())
		return nil, err
	}

	return new(emptypb.Empty), nil
}
