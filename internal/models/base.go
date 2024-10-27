package models

import (
	"database/sql"
	"time"
)

type Base struct {
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     sql.NullTime `json:"deleted_at"`
	ID            string       `json:"id"`
	TransactionID string       `json:"transaction_id"`
}
