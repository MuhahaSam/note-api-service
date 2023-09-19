package note_v1

import (
	"context"
	"fmt"

	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	fmt.Println("CreateNote")
	fmt.Println("title: ", req.GetTitle())
	fmt.Println("author: ", req.GetAuthor())
	fmt.Println("text: ", req.GetText())

	return &desc.CreateNoteResponse{
		Index: 1,
	}, nil

}
