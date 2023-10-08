package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/MuhahaSam/golangPractice/internal/app/db"
	"github.com/MuhahaSam/golangPractice/internal/app/entity"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	"github.com/google/uuid"
)

type NoteRepository struct {
	Repository
}

func (r *NoteRepository) Create(ctx context.Context, createNote *desc.CreateNoteRequest) (*uuid.UUID, error) {
	query, args, err := sq.Insert("note").
		PlaceholderFormat(sq.Dollar).
		Columns("author, title, text").
		Values(createNote.GetAuthor(), createNote.GetTitle(), createNote.GetText()).
		Suffix("returning id").
		ToSql()

	if err != nil {
		return nil, err
	}

	id, err := db.RunQueryToCreate[uuid.UUID](ctx, query, args)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (r *NoteRepository) Read(ctx context.Context, uuid uuid.UUID) (*entity.NoteEntity, error) {
	query, args, err := sq.Select("id, author, title, text").
		From("note").
		Where("id = $1", uuid).
		ToSql()
	if err != nil {
		return nil, err
	}

	note := &entity.NoteEntity{}
	err = db.RunQueryToGetFirst(note, query, args)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (e *NoteRepository) Update(id uuid.UUID, updateBody *desc.UpdateNoteBody) error {
	query, args, err := sq.Update("note").
		Set("author", updateBody.GetAuthor().Value).
		Set("title", updateBody.GetTitle().Value).
		Set("text", updateBody.GetText().Value).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	err = db.RunQuery(query, args)
	if err != nil {
		return err
	}

	return nil
}

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
