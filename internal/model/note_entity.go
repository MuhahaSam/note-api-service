package model

import (
	"database/sql"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type NoteUpdateBody struct {
	Author *wrapperspb.StringValue `db:"author"`
	Title  *wrapperspb.StringValue `db:"title"`
	Text   *wrapperspb.StringValue `db:"text"`
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
