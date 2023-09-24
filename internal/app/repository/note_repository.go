package repository

import (
	"github.com/MuhahaSam/golangPractice/internal/app/db"
	"github.com/MuhahaSam/golangPractice/internal/app/entity"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	"github.com/google/uuid"
)

type NoteRepository struct {
	Repository
}

func (r *NoteRepository) Create(createNote *desc.CreateNoteRequest) (uuid.UUID, error) {
	db := db.GetFakeDb()

	uuid := uuid.New()
	(*db)["Note"][uuid] = entity.NoteEntity{
		Id:     uuid,
		Title:  createNote.GetTitle(),
		Author: createNote.GetAuthor(),
		Text:   createNote.GetText(),
	}

	return uuid, nil
}

func (r *NoteRepository) Read(Id uuid.UUID) (entity.NoteEntity, error) {
	db := db.GetFakeDb()

	note := (*db)["Note"][Id]

	return note, nil
}

func (e *NoteRepository) Update(Id uuid.UUID, updateBody *desc.UpdateNoteBody) error {
	db := db.GetFakeDb()

	(*db)["Note"][Id] = entity.NoteEntity{
		Id:     Id,
		Author: updateBody.GetAuthor().GetValue(),
		Title:  updateBody.GetTitle().GetValue(),
		Text:   updateBody.GetText().GetValue(),
	}

	return nil
}

func (r *NoteRepository) Delete(Id uuid.UUID) error {
	db := db.GetFakeDb()

	noteContainer := (*db)["Note"]

	delete((*db)["Note"], Id)

	(*db)["Note"] = noteContainer

	return nil
}

var noteRepository *NoteRepository = nil

func GetNoteRepository() *NoteRepository {
	if noteRepository == nil {
		noteRepository = new(NoteRepository)
	}

	return noteRepository
}
