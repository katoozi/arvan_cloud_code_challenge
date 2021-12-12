package models

import "time"

type RedemptionRecord struct {
	Cellphone string    `json:"cellphone"`
	Timestamp time.Time `json:"timestamp"`
}

func NewRedemptionRecord(cellphone string) RedemptionRecord {
	return RedemptionRecord{
		Cellphone: cellphone,
		Timestamp: time.Now(),
	}
}
