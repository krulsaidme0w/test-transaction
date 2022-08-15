package processor

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Worker struct {
	db     *sql.DB
	userID string
}

func NewWorker(userID string, db *sql.DB) *Worker {
	return &Worker{
		db:     db,
		userID: userID,
	}
}

const (
	getTransactionAmountQuery = `
		SELECT key, amount FROM transaction
		WHERE user_id = $1 AND status = 'pending'
		ORDER BY created_at ASC
		LIMIT 1`

	getUserBalanceQuery = `
		SELECT balance FROM test_user
		WHERE id = $1`

	updateBalanceQuery = `
		UPDATE test_user
		SET balance = $1
		WHERE id = $2`

	updateTransactionStatusQuery = `
		UPDATE transaction
		SET status = $1
		WHERE key = $2`
)

func (w *Worker) Work() {
	ticker := time.NewTicker(100 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			w.processEarliestTransaction()
		}
	}
}

func (w *Worker) processEarliestTransaction() {
	transactionID, amount, err := w.getTransactionAmount()
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return
		default:
			log.Fatal(err)
		}
	}

	balance, err := w.getUserBalance()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(transactionID, amount, balance)

	if balance+amount <= 0 {
		w.updateTransactionStatus(transactionID, "failed")
		return
	}

	if err := w.updateTransactionStatus(transactionID, "status"); err != nil {
		log.Fatal(err)
	}

	if err := w.updateBalance(balance + amount); err != nil {
		log.Fatal(err)
	}

	fmt.Println("success")
}

func (w *Worker) getUserBalance() (int, error) {
	row := w.db.QueryRow(getUserBalanceQuery, w.userID)

	var amount int
	if err := row.Scan(&amount); err != nil {
		return 0, err
	}

	return amount, nil
}

func (w *Worker) getTransactionAmount() (string, int, error) {
	row := w.db.QueryRow(getTransactionAmountQuery, w.userID)

	var (
		balance int
		key     string
	)

	if err := row.Scan(&key, &balance); err != nil {
		return "", 0, err
	}

	return key, balance, nil
}

func (w *Worker) updateTransactionStatus(id, status string) error {
	_, err := w.db.Exec(updateTransactionStatusQuery, status, id)

	return err
}

func (w *Worker) updateBalance(balance int) error {
	_, err := w.db.Exec(updateBalanceQuery, balance, w.userID)

	return err
}
