package util

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DBCon *sql.DB

func ConnectDatabase() {
	var err error
	DBCon, err = sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = DBCon.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
