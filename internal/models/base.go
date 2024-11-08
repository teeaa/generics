package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	ID            *string        `json:"id"`
	TransactionID *string        `json:"transaction_id"`
}
