package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"test-transaction/internal/transactions/delivery/http"
	"test-transaction/internal/transactions/repository"
	"test-transaction/internal/transactions/usecase"
)

const (
	dbString = "host=localhost port=5432 user=user password=password dbname=test"
	addr     = ":8000"
)

func main() {
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := repository.NewUTransactionsRepository(db)
	usecase := usecase.NewTransactionsUseCase(repository)
	handler := http.NewHandler(usecase)

	router := fiber.New()
	group := router.Group("/test")
	group.Post("/transfer", handler.Transfer)

	if err := router.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
