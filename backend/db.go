package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func initDB() {
	var err error

	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=ecommerce_db sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database tidak konek:", err)
	}

	log.Println("Database Konek")
}
