package models

import "time"

type TransactionRequest struct {
	UserID string `json:"user_id"`
	Amount int64  `json:"amount"`
	Type   string `json:"type"`
}

type Transaction struct {
	UserID    string
	Amount    int64
	CreatedAt time.Time
	Status    string
	Key       string
}
