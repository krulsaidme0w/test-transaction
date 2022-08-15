package transactions

import (
	"context"
	"test-transaction/pkg/models"
)

type Repository interface {
	Deposit(ctx context.Context, userID string, amount int64) (string, error)
	Withdraw(ctx context.Context, userID string, amount int64) (string, error)
	GetTransaction(ctx context.Context, id string) (*models.Transaction, error)
}
