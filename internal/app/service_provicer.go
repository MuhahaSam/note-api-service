package app

import (
	"context"
	"log"

	"github.com/MuhahaSam/golangPractice/internal/config"
	"github.com/MuhahaSam/golangPractice/internal/db"
	"github.com/MuhahaSam/golangPractice/internal/repository"
	"github.com/MuhahaSam/golangPractice/internal/service/note"
	"github.com/vertica/vertica-sql-go/logger"
)

type serviceProvider struct {
	db         db.Client
	configPath string
	config     *config.NoteConfig
	logger     *logger.Logger

	// repositories
	noteRepository *repository.NoteRepository

	// services
	noteService *note.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

// GetDB ...
func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("can`t connect to db err: %s", err.Error())
		}
		s.db = dbc
	}

	return s.db
}

// GetConfig ...
func (s *serviceProvider) GetConfig() *config.NoteConfig {
	if s.config == nil {
		cfg, err := config.NewNoteConfig(s.configPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) GetLogger() *logger.Logger {
	if s.logger == nil {
		s.logger = logger.New("note logger")
	}

	return s.logger
}

// GetNoteRepository ...
func (s *serviceProvider) GetNoteRepository(ctx context.Context) repository.NoteRepository {
	if s.noteRepository == nil {
		s.noteRepository = repository.NewNoteRepository(s.GetDB(ctx))
	}

	return *s.noteRepository
}

// GetNoteService ...
func (s *serviceProvider) GetNoteService(ctx context.Context) *note.Service {
	if s.noteService == nil {
		s.noteService = note.NewService(
			s.GetLogger(),
			s.GetNoteRepository(ctx),
		)
	}

	return s.noteService
}
