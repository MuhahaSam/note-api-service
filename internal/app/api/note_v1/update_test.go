package note_v1

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/MuhahaSam/golangPractice/internal/model"
	noteMocks "github.com/MuhahaSam/golangPractice/internal/repository/mock"
	"github.com/MuhahaSam/golangPractice/internal/service/note"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	"github.com/go-faker/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdateNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		updatedText = faker.Username()
		uuid        = faker.UUIDDigit()

		repoErrorText = "error during note update"
		repoError     = errors.New(repoErrorText)

		updateBody = &model.NoteUpdateBody{
			Text: sql.NullString{String: updatedText, Valid: true},
		}

		validReq = &desc.UpdateRequest{
			Uuid: uuid,
			UpdateBody: &desc.UpdateBody{
				Text: &wrapperspb.StringValue{Value: updatedText},
			},
		}
	)

	noteMock := noteMocks.NewMockNoteRepositoryInterface(mockCtrl)

	api := NewNoteV1(note.NewService(noteMock))

	t.Run("success case", func(t *testing.T) {
		noteMock.EXPECT().Update(ctx, uuid, updateBody).Return(nil)
		_, err := api.UpdateNote(ctx, validReq)
		require.Nil(t, err)
	})

	t.Run("error case", func(t *testing.T) {
		noteMock.EXPECT().Update(ctx, uuid, updateBody).Return(repoError)
		_, err := api.UpdateNote(ctx, validReq)
		require.NotNil(t, err)
		require.Equal(t, repoErrorText, err.Error())
	})

}
