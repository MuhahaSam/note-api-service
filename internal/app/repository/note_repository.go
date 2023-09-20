package repository

import (
	"github.com/MuhahaSam/golangPractice/internal/app/db"
	"github.com/MuhahaSam/golangPractice/internal/app/entity"

	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

type NoteRepository struct {
	Repository
}

func (r *NoteRepository) Create(createNote *desc.CreateNoteRequest) (int64, error) {
	db := db.GetFakeDb()

	index := int64(len((*db)["Note"]))
	(*db)["Note"] = append((*db)["Note"], entity.NoteEntity{
		Index:  index,
		Title:  createNote.GetTitle(),
		Author: createNote.GetAuthor(),
		Text:   createNote.GetText(),
	})

	return index, nil
}

var noteRepository *NoteRepository = nil

func GetNoteRepository() *NoteRepository {
	if noteRepository == nil {
		noteRepository = new(NoteRepository)
	}

	return noteRepository
}
