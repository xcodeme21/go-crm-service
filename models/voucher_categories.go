package models

import (
	"time"
)

type VoucherCategories struct {
	ID           int       `json:"id"`
	Name  string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
