package converter

import (
	"database/sql"

	"github.com/MuhahaSam/golangPractice/internal/model"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

func ToNoteModel(createRequest *desc.CreateRequest) *model.NoteEntity {
	return &model.NoteEntity{
		NoteInfo: model.NoteInfo{
			Title:  createRequest.GetTitle(),
			Text:   createRequest.GetText(),
			Author: createRequest.GetAuthor(),
		},
	}
}

func ToNoteUpdateBody(updateBody *desc.UpdateBody) *model.NoteUpdateBody {
	var modelUpdateBody = model.NoteUpdateBody{}

	if updateBody.GetAuthor() != nil {
		modelUpdateBody.Author = sql.NullString{String: updateBody.GetAuthor().Value, Valid: true}
	}
	if updateBody.GetTitle() != nil {
		modelUpdateBody.Title = sql.NullString{String: updateBody.GetTitle().Value, Valid: true}
	}
	if updateBody.GetText() != nil {
		modelUpdateBody.Text = sql.NullString{String: updateBody.GetText().Value, Valid: true}
	}
	return &modelUpdateBody
}
