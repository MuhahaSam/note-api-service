package db

import (
	"context"
	"fmt"
	"log"

	"github.com/MuhahaSam/golangPractice/config"
	"github.com/MuhahaSam/golangPractice/internal/app/entity"
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

var dbModule *DbModule = &DbModule{}

func GetDbModule() *DbModule {
	return dbModule
}

func RunQueryToCreate[T any](ctx context.Context, query string, args []any) (*T, error) {
	rows, err := dbModule.DbConnection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id T
	rows.Next()
	err = rows.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

/*
Я не знаю насколько правильно такое решение.
Подозревю, что это мега не красиво. Если будут идеи
как сделать хорошо прошу дать фидбек
*/
func RunQueryToGetFirst[T entity.Entity](object *T, query string, args []any) error {
	rows, err := dbModule.DbConnection.Queryx(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	rows.Next()
	err = rows.StructScan(&object)
	if err != nil {
		return err
	}

	return nil
}
