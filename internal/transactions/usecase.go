package transactions

import (
	"context"
)

type UseCase interface {
	Deposit(ctx context.Context, userID string, amount int64) error
	Withdraw(ctx context.Context, userID string, amount int64) error
}
