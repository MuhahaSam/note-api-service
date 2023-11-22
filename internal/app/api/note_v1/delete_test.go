package note_v1

import (
	"context"
	"errors"
	"testing"

	noteMocks "github.com/MuhahaSam/golangPractice/internal/repository/mock"
	"github.com/MuhahaSam/golangPractice/internal/service/note"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	"github.com/go-faker/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDeleteNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		uuid = faker.UUIDDigit()

		repoErrorText = "error during note update"
		repoError     = errors.New(repoErrorText)

		validReq = &desc.DeleteRequest{
			Uuid: uuid,
		}
	)

	noteMock := noteMocks.NewMockNoteRepositoryInterface(mockCtrl)

	api := NewNoteV1(note.NewService(noteMock))

	t.Run("success case", func(t *testing.T) {
		noteMock.EXPECT().Delete(ctx, uuid).Return(nil)
		_, err := api.DeleteNote(ctx, validReq)
		require.Nil(t, err)
	})

	t.Run("error case", func(t *testing.T) {
		noteMock.EXPECT().Delete(ctx, uuid).Return(repoError)
		_, err := api.DeleteNote(ctx, validReq)
		require.NotNil(t, err)
		require.Equal(t, repoErrorText, err.Error())
	})

}
