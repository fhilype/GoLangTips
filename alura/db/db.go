package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err == nil {
		return db
	} else {
		panic(err.Error())
	}
}
