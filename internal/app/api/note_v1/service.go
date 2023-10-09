package note_v1

import (
	"github.com/MuhahaSam/golangPractice/config"
	db "github.com/MuhahaSam/golangPractice/internal/app/db"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

type Note struct {
	desc.NoteServiceServer
}

func (n *Note) Init() error {
	err := db.GetDbModule().Open(&config.GetConfig().DbConfig)
	if err != nil {
		return err
	}
	return nil
}

func (n *Note) Destructor() {
	db.GetDbModule().Close()
}

func NewNote() *Note {
	return &Note{}
}
