package repository

import (
	"context"
	"time"

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

	rows, err := db.GetDbModule().DbConnection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id uuid.UUID
	rows.Next()
	err = rows.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *NoteRepository) Read(ctx context.Context, uuid uuid.UUID) (*entity.NoteEntity, error) {
	query, args, err := sq.Select("id, author, title, text").
		From("note").
		Where("id = $1 and deleted_at is null", uuid).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.GetDbModule().DbConnection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	note := &entity.NoteEntity{}
	rows.Next()
	err = rows.Scan(&note.Id, &note.Author, &note.Title, &note.Text)
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
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	rows, err := db.GetDbModule().DbConnection.Queryx(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *NoteRepository) Delete(id uuid.UUID) error {
	query, args, err := sq.Update("note").
		Set("deleted_at", time.Now()).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	rows, err := db.GetDbModule().DbConnection.Queryx(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

var noteRepository *NoteRepository = nil

func GetNoteRepository() *NoteRepository {
	if noteRepository == nil {
		noteRepository = new(NoteRepository)
	}

	return noteRepository
}
