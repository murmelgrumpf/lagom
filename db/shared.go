package db

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func GetDb() *sql.DB {
	if db == nil {
		var error error
		db, error = sql.Open("sqlite3", os.Getenv("DB_LOCATION"))

		if error != nil {
			panic(error)
		}
	}
	return db
}

func Prepare(query string) (*sql.Stmt, error) {
	statement, error := GetDb().Prepare(query)

	if error != nil {
		slog.Error("An error occured while preparing query '"+query+"'", slog.Any("error", error))
	}

	return statement, error
}

func PreparePanic(query string) *sql.Stmt {
	statement, error := Prepare(query)
	if error != nil {
		panic(error)
	}
	return statement
}

func Query(query string) (*sql.Rows, error) {
	rows, error := GetDb().Query(query)

	if error != nil {
		slog.Error("An error occured while executing query '"+query+"'", slog.Any("error", error))
	}
	return rows, error
}

func QueryPanic(query string) *sql.Rows {
	rows, error := Query(query)
	if error != nil {
		panic(error)
	}
	return rows
}
