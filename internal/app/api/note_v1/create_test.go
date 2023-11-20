package note_v1

import (
	"context"
	"errors"
	"testing"

	"github.com/MuhahaSam/golangPractice/internal/model"
	noteMocks "github.com/MuhahaSam/golangPractice/internal/repository/mock"
	"github.com/MuhahaSam/golangPractice/internal/service/note"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	"github.com/go-faker/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		title  = faker.Name()
		text   = faker.Username()
		author = faker.ChineseFirstName()
		uuid   = faker.UUIDDigit()

		repoErrorText = "error during note create"

		repoError = errors.New(repoErrorText)

		validReq = &desc.CreateRequest{
			Title:  title,
			Text:   text,
			Author: author,
		}

		validNoteEntity = &model.NoteEntity{
			NoteInfo: model.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
		}

		validRes = &desc.CreateResponse{
			Uuid: uuid,
		}
	)

	noteMock := noteMocks.NewMockNoteRepositoryInterface(mockCtrl)

	api := NewNoteV1(note.NewService(noteMock))

	t.Run("success case", func(t *testing.T) {
		noteMock.EXPECT().Create(ctx, validNoteEntity).Return(&uuid, nil)
		res, err := api.CreateNote(ctx, validReq)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("error case", func(t *testing.T) {
		noteMock.EXPECT().Create(ctx, validNoteEntity).Return(nil, repoError)
		_, err := api.CreateNote(ctx, validReq)
		require.NotNil(t, err)
		require.Equal(t, repoErrorText, err.Error())
	})

}
