package repository

//go:generate mockgen --build_flags=--mod=mod -destination=mock/mock_note_repository.go -package=mocks . NoteRepositoryInterface
import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/MuhahaSam/golangPractice/internal/db"
	"github.com/MuhahaSam/golangPractice/internal/model"
)

const (
	EntityName = "note"
	Id         = "id"
	Author     = "author"
	Title      = "title"
	Text       = "text"
	CreatedAt  = "created_at"
	UpdatedAt  = "updated_at"
	DeletedAt  = "deleted_at"
)

type NoteRepositoryInterface interface {
	Create(ctx context.Context, note *model.NoteEntity) (*string, error)
	Get(ctx context.Context, uuid string) (*model.NoteEntity, error)
	Update(ctx context.Context, uuid string, noteInfo *model.NoteUpdateBody) error
	Delete(ctx context.Context, uuid string) error
}
type NoteRepository struct {
	db db.Client
}

func (r *NoteRepository) Create(ctx context.Context, note *model.NoteEntity) (*string, error) {
	query, args, err := sq.Insert(EntityName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			Author,
			Title,
			Text,
		).
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

	var id string
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
	}

	return &id, nil
}

func (r *NoteRepository) Get(ctx context.Context, uuid string) (*model.NoteEntity, error) {
	query, args, err := sq.
		Select(
			Id,
			Author,
			Title,
			Text).
		From(EntityName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{Id: uuid, DeletedAt: nil}).
		ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "repository.GetNote",
		QueryRaw: query,
	}

	var notes []*model.NoteEntity

	err = r.db.DB().SelectContext(ctx, &notes, q, args...)
	if err != nil {
		return nil, err
	}

	if len(notes) == 0 {
		return nil, errors.New("note not found")
	}

	return notes[0], nil
}

func (r *NoteRepository) Update(ctx context.Context, uuid string, noteInfo *model.NoteUpdateBody) error {
	builder := sq.Update(EntityName).
		Where(sq.Eq{Id: uuid}).
		Set(UpdatedAt, time.Now()).
		PlaceholderFormat(sq.Dollar)

	if noteInfo.Author.Valid {
		builder = builder.Set(Author, noteInfo.Author.String)
	}
	if noteInfo.Title.Valid {
		builder = builder.Set(Title, noteInfo.Title.String)
	}
	if noteInfo.Text.Valid {
		builder = builder.Set(Text, noteInfo.Text.String)
	}

	query, args, err := builder.ToSql()
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
	query, args, err := sq.Update(EntityName).
		Set(DeletedAt, time.Now()).
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

func NewNoteRepository(db db.Client) *NoteRepository {
	return &NoteRepository{
		db: db,
	}
}
