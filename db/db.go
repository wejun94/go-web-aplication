package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	connStr := "user=postgres dbname=store password=Wejun@_!123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}
