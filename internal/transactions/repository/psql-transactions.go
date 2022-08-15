package repository

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"strconv"
	"time"

	"test-transaction/pkg/models"
)

type transactionsRepository struct {
	DB *sql.DB
}

func NewTransactionsRepository(db *sql.DB) *transactionsRepository {
	return &transactionsRepository{
		DB: db,
	}
}

const (
	transactionQuery = `
		INSERT INTO transaction(user_id, amount, created_at, status, key) 
		VALUES ($1, $2, $3, $4, $5)`

	getTransactionQuery = `
		SELECT user_id, amount, created_at, status
		FROM transaction 
		WHERE key = $1`
)

func (r *transactionsRepository) Deposit(ctx context.Context, userID string, amount int64) (string, error) {
	transactionID := hash(userID + strconv.Itoa(int(amount)) + time.Now().String())

	_, err := r.DB.Exec(transactionQuery, userID, amount, time.Now(), "pending", transactionID)

	return transactionID, err
}

func (r *transactionsRepository) Withdraw(ctx context.Context, userID string, amount int64) (string, error) {
	transactionID := hash(userID + strconv.Itoa(int(amount)) + time.Now().String())

	_, err := r.DB.Exec(transactionQuery, userID, -1*amount, time.Now(), "pending", transactionID)

	return transactionID, err
}

func (r *transactionsRepository) GetTransaction(ctx context.Context, id string) (*models.Transaction, error) {
	row := r.DB.QueryRow(getTransactionQuery, id)

	transaction := &models.Transaction{}
	if err := row.Scan(&transaction.UserID, &transaction.Amount, &transaction.CreatedAt, &transaction.Status); err != nil {
		return nil, err
	}

	return transaction, nil
}

func hash(str string) string {
	hash := sha256.Sum256([]byte(str))

	return hex.EncodeToString(hash[:])
}
