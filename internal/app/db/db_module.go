package db

import (
	"fmt"
	"sync"

	"github.com/google/uuid"

	"github.com/MuhahaSam/golangPractice/internal/app/entity"
)

type DbModuleInterface interface {
	CreateConnection(dsn string) error
	Connect() error
	Close() error
}

type FakeDbModule struct {
	DbModuleInterface
}

func (db *FakeDbModule) Connect(dsn string) error {
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

type FakeDb struct {
	mu      sync.RWMutex
	records map[string]map[uuid.UUID]entity.NoteEntity
}

func (f *FakeDb) Read(uuid uuid.UUID) entity.NoteEntity {
	f.mu.RLock()
	defer f.mu.Unlock()
	return f.records["Note"][uuid]
}

func (f *FakeDb) Write(uuid uuid.UUID, note entity.NoteEntity) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.records["Note"][uuid] = note
}

func (f *FakeDb) Delete(uuid uuid.UUID) {
	f.mu.Lock()
	defer f.mu.Unlock()
	delete(f.records["Note"], uuid)
}

var fakeDb = FakeDb{
	records: make(map[string]map[uuid.UUID]entity.NoteEntity),
}

func GetFakeDb() *FakeDb {
	if fakeDb.records["Note"] == nil {
		fakeDb.records["Note"] = make(map[uuid.UUID]entity.NoteEntity)
	}
	return &fakeDb
}
