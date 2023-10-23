package models

import (
	"time"
)

type Vouchers struct {
	ID        int       `json:"id"`
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

type VouchersFindAllResponse struct {
	ID        int       `json:"id"`
	VoucherName   string `json:"voucher_name"`
	SeriesId   int `json:"series_id"`
	IsExternalVoucher bool   `json:"is_external_voucher"`
	CampaignId   int `json:"campaign_id"`
	CampaignVoucherId   int `json:"campaign_voucher_id"`
	Point   int64 `json:"point"`
	Tier   string `json:"tier"`
	CategoryId   int `json:"category_id"`
	CategoryName         string `json:"category_name"`
	IsLimited bool   `json:"is_limited"`
	StartDate         time.Time `json:"start_date"`
    EndDate           time.Time `json:"end_date"`
	Image   string `json:"image"`
	Description   string `json:"description"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FilterVouchers struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	SortBy  string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
	Start   string `json:"start_date"`
	End     string `json:"end_date"`
	Status  string `json:"status"`
	Search  string `json:"search"`
}