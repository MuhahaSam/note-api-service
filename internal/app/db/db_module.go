package db

import (
	"log"

	"github.com/MuhahaSam/golangPractice/config"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type DbModuleInterface interface {
	CreateConnection(dsn string) error
	Open()
	Close()
	GetDb()
}

type DbModule struct {
	DbModuleInterface
	DbConnection *sqlx.DB
}

func (d *DbModule) Open(dbConfig *config.DbConfig) error {
	db, err := sqlx.Open(dbConfig.Engine, dbConfig.Dsn)
	if err != nil {
		log.Printf("error during connection to data base: %s", err)
		return err
	}

	d.DbConnection = db
	return nil
}

func (d *DbModule) Close() {
	defer d.DbConnection.Close()
}

var dbModule *DbModule = &DbModule{}

func GetDbModule() *DbModule {
	return dbModule
}
