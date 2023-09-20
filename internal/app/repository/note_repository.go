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

func (e *NoteRepository) Read(index int64) (entity.NoteEntity, error) {
	db := db.GetFakeDb()

	note := (*db)["Note"][index]

	return note, nil
}

func (e *NoteRepository) Update(index int64, updateBody *desc.UpdateNoteBody) error {
	db := db.GetFakeDb()

	(*db)["Note"][index] = entity.NoteEntity{
		Index:  index,
		Author: updateBody.GetAuthor(),
		Title:  updateBody.GetTitle(),
		Text:   updateBody.GetText(),
	}

	return nil
}

var noteRepository *NoteRepository = nil

func GetNoteRepository() *NoteRepository {
	if noteRepository == nil {
		noteRepository = new(NoteRepository)
	}

	return noteRepository
}
