package transactions

import (
	"context"
)

type Repository interface {
	Deposit(ctx context.Context, userID string, amount int64) (string, error)
	Withdraw(ctx context.Context, userID string, amount int64) (string, error)
}
