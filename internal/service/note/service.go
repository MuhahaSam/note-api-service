package note

import (
	"github.com/MuhahaSam/golangPractice/internal/repository"
)

// Service ...
type Service struct {
	noteRepository repository.NoteRepositoryInterface
}

// NewService ...
func NewService(
	noteRepository repository.NoteRepositoryInterface,
) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}
