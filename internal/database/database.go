package database

import (
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

func ConnectToDB() {
	var err error
	DB, err = sqlx.Connect("postgres", "user=postgres password=3848344 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
