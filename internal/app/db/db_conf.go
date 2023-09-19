package db

import (
	"fmt"
)

type DbConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
}

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

func (db *FakeDbModule) Close(config DbConfig) error {
	fmt.Println("close data base connection")
	return nil
}

var fakeDbConfig = DbConfig{
	Host:     "postgres.host",
	User:     "user",
	Password: "password",
	DbName:   "dbName",
}
var fakeDbModule *FakeDbModule = nil

func GetDbModuleInstance() *FakeDbModule {
	if fakeDbModule == nil {
		fakeDbModule = new(FakeDbModule)
		fakeDbModule.CreateConnection(fakeDbConfig)
	}

	return fakeDbModule
}
