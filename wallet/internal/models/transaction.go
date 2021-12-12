package models

import "time"

// Transaction is the transaction structure in the wallet service
type Transaction struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
}
