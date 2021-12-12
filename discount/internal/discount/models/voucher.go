package models

import "time"

// Voucher is the voucher model
type Voucher struct {
	Code          string    `json:"voucher"`
	DiscountPrice uint64    `json:"discount_price"`
	MaxUseTime    uint64    `json:"max_use_time"`
	CreatedAt     time.Time `json:"created_at"`
}
