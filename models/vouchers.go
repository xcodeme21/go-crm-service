package models

import (
	"time"
)

type VoucherCategories struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	VoucherName   string `json:"voucher_name"`
	SeriesId   int `json:"series_id"`
	IsExternalVoucher bool   `json:"is_external_voucher"`
	CampaignId   int `json:"campaign_id"`
	CampaignVoucherId   int `json:"campaign_voucher_id"`
	Point   int64 `json:"point"`
	Tier   string `json:"tier"`
	CategoryId   int `json:"category_id"`
	IsLimited bool   `json:"is_limited"`
	StartDate         time.Time `json:"start_date"`
    EndDate           time.Time `json:"end_date"`
	Image   string `json:"image"`
	Description   string `json:"description"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}