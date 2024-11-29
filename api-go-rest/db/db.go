package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connection := "user=postgres dbname=api-go-rest password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err == nil {
		return db
	} else {
		panic(err.Error())
	}
}
