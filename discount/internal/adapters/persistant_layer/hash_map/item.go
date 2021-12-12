package hashmap

import (
	"sync"
	"time"
)

type storageItem struct {
	redemptionRecords sync.Map
	maxUseTime        uint64
	totalMaxUsedTime  uint64
	discountPrice     uint64
	createdAt         time.Time
}

// newStorageItem will create a new storage item.
func newStorageItem(createdAt time.Time, discountPrice, maxUseTime, totalMaxUsedTime uint64) *storageItem {
	return &storageItem{
		discountPrice:     discountPrice,
		createdAt:         createdAt,
		maxUseTime:        maxUseTime,
		totalMaxUsedTime:  totalMaxUsedTime,
		redemptionRecords: sync.Map{},
	}
}
