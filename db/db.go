package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Conexão com banco de dados

func ConnectionDataBase() *sql.DB {

	connection := "user=postgres dbname=jow_loja password=dev host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}
