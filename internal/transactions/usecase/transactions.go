package usecase

import (
	"context"

	"test-transaction/internal/transactions"
)

type transactionsUseCase struct {
	repository transactions.Repository
}

func NewTransactionsUseCase(repository transactions.Repository) *transactionsUseCase {
	return &transactionsUseCase{
		repository: repository,
	}
}

func (u *transactionsUseCase) Withdraw(ctx context.Context, userID string, amount int64) error {
	return u.repository.Withdraw(ctx, userID, amount)
}

func (u *transactionsUseCase) Deposit(ctx context.Context, userID string, amount int64) error {
	return u.repository.Deposit(ctx, userID, amount)
}