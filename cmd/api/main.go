package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"

	"test-transaction/internal/transactions/delivery/http"
	"test-transaction/internal/transactions/repository"
	"test-transaction/internal/transactions/usecase"
)

const (
	dbString = "host=0.0.0.0 port=5432 user=user password=password dbname=test sslmode=disable"
	addr     = ":8000"
)

func main() {
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := repository.NewTransactionsRepository(db)
	usecase := usecase.NewTransactionsUseCase(repository)
	handler := http.NewHandler(usecase)

	router := fiber.New()
	router.Use(logger.New())
	router.Post("/transfer", handler.Transfer)
	router.Get("/transaction/:id", handler.GetTransaction)

	if err := router.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
