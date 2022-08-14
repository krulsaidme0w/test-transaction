package repository

import (
	"context"
	"database/sql"
	"time"
)

type transactionsRepository struct {
	DB *sql.DB
}

func NewUTransactionsRepository(db *sql.DB) *transactionsRepository {
	return &transactionsRepository{
		DB: db,
	}
}

const (
	transactionQuery = `
		INSERT INTO transactions(user_id, amount, created_at) 
		VALUES ($1, $2, $3)`
)

func (r *transactionsRepository) Deposit(ctx context.Context, userID string, amount int64) error {
	_, err := r.DB.Exec(transactionQuery, userID, amount, time.Now())

	return err
}

func (r *transactionsRepository) Withdraw(ctx context.Context, userID string, amount int64) error {

	_, err := r.DB.Exec(transactionQuery, userID, -1*amount, time.Now())

	return err
}
