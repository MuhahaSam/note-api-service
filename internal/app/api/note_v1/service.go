package note_v1

import (
	"github.com/MuhahaSam/golangPractice/config"
	db "github.com/MuhahaSam/golangPractice/internal/app/db"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
)

type Note struct {
	desc.NoteServiceServer
}

func (n *Note) New() *Note {

	db.GetDbModuleInstance().Open(&config.GetConfig().DbConfig)
	return n
}

func (n *Note) Destructor() {
	db.GetDbModuleInstance().Close()
}

func NewNote() *Note {
	return &Note{}
}
