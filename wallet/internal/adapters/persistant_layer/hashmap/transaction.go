package hashmap

import (
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/katoozi/avran_cloud_code_challenge_wallet_service/internal/models"
)

func (hm *hashMap) Transaction(cellphone string, amount int64, timestamp time.Time, description string) (models.Transaction, int64, error) {
	item, ok := hm.users.Load(cellphone)
	if !ok {
		item = new(models.Wallet)
		hm.users.Store(cellphone, item)
	}

	wallet := item.(*models.Wallet)

	id := hm.generateID()

	atomic.AddInt64(&wallet.CurrentBalance, amount)

	newTRX := &models.Transaction{
		ID:          id,
		CreatedAt:   timestamp,
		Amount:      amount,
		Description: description,
	}

	wallet.Transactions.Store(id, newTRX)

	return *newTRX, wallet.CurrentBalance, nil
}

func (hm *hashMap) generateID() string {
	// TODO: put limit on for loop iterations number
	for {
		id := uuid.New().String()
		_, ok := hm.transactionIDs.LoadOrStore(id, nil)
		if !ok {
			return id
		}
	}
}
