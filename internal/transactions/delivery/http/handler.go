package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	"test-transaction/internal/transactions"
	"test-transaction/pkg/models"
)

type TransactionsHandler struct {
	usecase transactions.UseCase
}

func NewHandler(usecase transactions.UseCase) *TransactionsHandler {
	return &TransactionsHandler{
		usecase: usecase,
	}
}

func (h *TransactionsHandler) Transfer(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	fmt.Println("asd")

	transactionRequest := new(models.TransactionRequest)

	err := json.Unmarshal(c.Body(), &transactionRequest)
	if err != nil {
		c.Context().Error(err.Error(), http.StatusBadRequest)
		body, _ := json.Marshal(err.Error())
		c.Context().SetBody(body)
		return nil
	}

	transactionID := ""
	switch transactionRequest.Type {
	case "withdraw":
		transactionID, err = h.usecase.Withdraw(ctx, transactionRequest.UserID, transactionRequest.Amount)
	case "deposit":
		transactionID, err = h.usecase.Deposit(ctx, transactionRequest.UserID, transactionRequest.Amount)
	}
	if err != nil {
		c.Context().Error(err.Error(), http.StatusBadRequest)
		body, _ := json.Marshal(err.Error())
		c.Context().SetBody(body)
		return nil
	}

	body, _ := json.Marshal(transactionID)
	c.Context().Success("application:json", body)

	return nil
}
