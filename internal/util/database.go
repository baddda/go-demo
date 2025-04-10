package util

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDatabase() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
