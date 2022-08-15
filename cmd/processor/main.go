package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"test-transaction/internal/processor"
)

const (
	dbString = "host=0.0.0.0 port=5432 user=user password=password dbname=test sslmode=disable"
)

func main() {
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	p := processor.NewProcessor(db)

	p.Run()
}
