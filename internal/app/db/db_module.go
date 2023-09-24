package db

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/MuhahaSam/golangPractice/internal/app/entity"
)

type DbModuleInterface interface {
	CreateConnection(config DbConfig) error
	Connect() error
	Close() error
}

type FakeDbModule struct {
	DbModuleInterface
}

func (db *FakeDbModule) Connect(config *DbConfig) error {
	fmt.Println("connect to data base")
	return nil
}

func (db *FakeDbModule) Close() error {
	fmt.Println("close data base connection")
	return nil
}

var fakeDbModule *FakeDbModule = nil

func GetDbModuleInstance() *FakeDbModule {
	if fakeDbModule == nil {
		fakeDbModule = new(FakeDbModule)
	}

	return fakeDbModule
}

var fakeDb map[string]map[uuid.UUID]entity.NoteEntity = nil

func GetFakeDb() *map[string]map[uuid.UUID]entity.NoteEntity {
	if fakeDb == nil {
		fakeDb = make(map[string]map[uuid.UUID]entity.NoteEntity)
		fakeDb["Note"] = make(map[uuid.UUID]entity.NoteEntity)
	}
	return &fakeDb
}
