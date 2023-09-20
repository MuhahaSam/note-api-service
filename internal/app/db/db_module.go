package db

import (
	"fmt"

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

func (db *FakeDbModule) CreateConnection(config DbConfig) error {
	fmt.Println("data base connection created")
	return nil
}

func (db *FakeDbModule) Connect() error {
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
		fakeDbModule.CreateConnection(fakeDbConfig)
	}

	return fakeDbModule
}

type FakeEntityContainer struct {
	Index   int64
	Records map[int64]entity.NoteEntity
}

var fakeDb *map[string]FakeEntityContainer = nil

func GetFakeDb() *map[string]FakeEntityContainer {
	if fakeDb == nil {
		fakeDb = &map[string]FakeEntityContainer{
			"Note": FakeEntityContainer{
				Index:   0,
				Records: map[int64]entity.NoteEntity{},
			},
		}
	}
	return fakeDb
}
