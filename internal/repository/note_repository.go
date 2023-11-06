package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/MuhahaSam/golangPractice/internal/db"
	"github.com/MuhahaSam/golangPractice/internal/model"
	"github.com/google/uuid"
)

type NoteRepository struct {
	db db.Client
}

// NewEventRepository ...
func NewNoteRepository(db db.Client) *NoteRepository {
	return &NoteRepository{
		db: db,
	}
}

func (r *NoteRepository) Create(ctx context.Context, note *model.NoteEntity) (*uuid.UUID, error) {
	query, args, err := sq.Insert("note").
		PlaceholderFormat(sq.Dollar).
		Columns("author, title, text").
		Values(note.NoteInfo.Author, note.NoteInfo.Title, note.NoteInfo.Text).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "repository.CreateNote",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id uuid.UUID
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
	}

	return &id, nil
}

func (r *NoteRepository) Get(ctx context.Context, uuid string) (*model.NoteEntity, error) {
	query, args, err := sq.Select("id, author, title, text").
		From("note").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": uuid, "deleted_at": nil}).
		ToSql()

	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "repository.GetNote",
		QueryRaw: query,
	}

	var notes []*model.NoteEntity = []*model.NoteEntity{}

	err = r.db.DB().SelectContext(ctx, &notes, q, args...)
	if err != nil {
		return nil, err
	}

	return &*notes[0], nil
}

func (r *NoteRepository) Update(ctx context.Context, uuid string, noteInfo *model.NoteUpdateBody) error {
	updateInfoMap := make(map[string]interface{})

	if noteInfo.Author.ProtoReflect().IsValid() {
		updateInfoMap["author"] = noteInfo.Author.GetValue()
	}
	if noteInfo.Text.ProtoReflect().IsValid() {
		updateInfoMap["text"] = noteInfo.Text.GetValue()
	}
	if noteInfo.Title.ProtoReflect().IsValid() {
		updateInfoMap["title"] = noteInfo.Title.GetValue()
	}

	query, args, err := sq.Update("note").
		SetMap(updateInfoMap).
		Where(sq.Eq{"id": uuid}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "repository.UpdateNote",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *NoteRepository) Delete(ctx context.Context, uuid string) error {
	query, args, err := sq.Update("note").
		Set("deleted_at", time.Now()).
		Where(sq.Eq{"id": uuid}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "repository.DeleteNote",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
