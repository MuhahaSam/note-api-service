package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/MuhahaSam/golangPractice/internal/app/db"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	"github.com/google/uuid"
)

type NoteRepository struct {
	Repository
}

func (r *NoteRepository) Create(ctx context.Context, createNote *desc.CreateNoteRequest) (*uuid.UUID, error) {
	builder := sq.Insert("note").
		PlaceholderFormat(sq.Dollar).
		Columns("author, title, text").
		Values(createNote.GetAuthor(), createNote.GetTitle(), createNote.GetText()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.
		GetDbModuleInstance().
		GetDbConnection().
		QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id string
	rows.Next()
	err = rows.Scan(&id)
	if err != nil {
		return nil, err
	}

	note_uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return &note_uuid, nil
}

// func (r *NoteRepository) Read(id uuid.UUID) (entity.NoteEntity, error) {
// 	db := db.GetFakeDb()
// 	note := db.Read(id)
// 	return note, nil
// }

// func (e *NoteRepository) Update(id uuid.UUID, updateBody *desc.UpdateNoteBody) error {
// 	db := db.GetFakeDb()

// 	db.Write(id, entity.NoteEntity{
// 		Id:     id,
// 		Author: updateBody.GetAuthor().GetValue(),
// 		Title:  updateBody.GetTitle().GetValue(),
// 		Text:   updateBody.GetText().GetValue(),
// 	})

// 	return nil
// }

// func (r *NoteRepository) Delete(id uuid.UUID) error {
// 	db := db.GetFakeDb()
// 	db.Delete(id)
// 	return nil
// }

var noteRepository *NoteRepository = nil

func GetNoteRepository() *NoteRepository {
	if noteRepository == nil {
		noteRepository = new(NoteRepository)
	}

	return noteRepository
}
