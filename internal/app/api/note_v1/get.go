package note_v1

import (
	"context"
	"log"

	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

func (n *Implementation) GetNote(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	note, err := n.noteService.Get(ctx, req)
	if err != nil {
		log.Printf("error while reading note: %s", err.Error())
		return nil, err
	}

	return &desc.GetResponse{
		Uuid:   note.Id,
		Title:  note.NoteInfo.Title,
		Author: note.NoteInfo.Author,
		Text:   note.NoteInfo.Text,
	}, nil
}
