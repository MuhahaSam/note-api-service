package note

import (
	"github.com/MuhahaSam/golangPractice/internal/repository"
	"github.com/vertica/vertica-sql-go/logger"
)

// Service ...
type Service struct {
	logger         *logger.Logger
	noteRepository repository.NoteRepository
}

// NewService ...
func NewService(
	logger *logger.Logger,
	noteRepository repository.NoteRepository,
) *Service {
	return &Service{
		logger:         logger,
		noteRepository: noteRepository,
	}
}
