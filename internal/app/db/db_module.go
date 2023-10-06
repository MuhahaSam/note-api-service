package db

import (
	"fmt"
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

func (d *DbModule) Open(dbConfig *config.DbConfig) {
	dbDns := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.User, dbConfig.Password)

	db, err := sqlx.Open(dbConfig.Engine, dbDns)
	if err != nil {
		log.Fatalf("error during connection to data base: %s", err)
	}

	d.DbConnection = db
}

func (d *DbModule) Close() {
	defer d.DbConnection.Close()
}

func (d *DbModule) GetDbConnection() *sqlx.DB {
	return d.DbConnection
}

var dbModule *DbModule = &DbModule{}

func GetDbModuleInstance() *DbModule {
	return dbModule
}
