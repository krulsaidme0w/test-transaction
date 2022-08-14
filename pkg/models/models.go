package models

type TransactionRequest struct {
	UserID string `json:"user_id"`
	Amount int64  `json:"amount"`
	Type   string `json:"type"`
}
