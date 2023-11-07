package model

import (
	"database/sql"
)

type NoteUpdateBody struct {
	Author sql.NullString `db:"author"`
	Title  sql.NullString `db:"title"`
	Text   sql.NullString `db:"text"`
}

type NoteInfo struct {
	Author string `db:"author"`
	Title  string `db:"title"`
	Text   string `db:"text"`
}

type NoteEntity struct {
	Id        string       `db:"id"`
	NoteInfo  NoteInfo     `db:""`
	CreatedAt sql.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
}
