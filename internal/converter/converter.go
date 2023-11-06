package converter

import (
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
	return &model.NoteUpdateBody{
		Title:  updateBody.GetTitle(),
		Author: updateBody.GetAuthor(),
		Text:   updateBody.GetText(),
	}
}
