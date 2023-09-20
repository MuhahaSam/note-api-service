package repository

import (
	db "github.com/MuhahaSam/golangPractice/internal/app/db"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

type NoteRepository struct {
	Repository
}

func (r *NoteRepository) Create(createNote *desc.CreateNoteRequest) int {
	db := db.GetFakeDb()
	(*db)["Note"] = append((*db)["Note"], createNote)

	return len((*db)["Note"])
}
