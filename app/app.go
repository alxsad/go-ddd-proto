package app

import (
	"app/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"os"
)

func NewZeroLogger() zerolog.Logger {
	return zerolog.New(os.Stdout).With().Timestamp().Stack().Logger()
}

func NewDb() *sqlx.DB {
	var db, err = sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		panic(fmt.Errorf("can not connect to db: %w", err))
	}

	schema := `CREATE TABLE users (id text, name text, email text)`

	_, err = db.Exec(schema)
	if err != nil {
		panic(fmt.Errorf("can not create schema: %w", err))
	}

	return db
}

type Application struct {
	log   zerolog.Logger
	tm    TransactionManager
	users domain.UsersService
}

func NewApplication(log zerolog.Logger, tm TransactionManager, users domain.UsersService) Application {
	return Application{
		log:   log,
		tm:    tm,
		users: users,
	}
}
