package entity

type Entity interface {
}

type NoteEntity struct {
	Entity
	Id     string
	Author string
	Title  string
	Text   string
}
