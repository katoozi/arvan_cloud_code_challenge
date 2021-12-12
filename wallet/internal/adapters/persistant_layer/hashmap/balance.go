package hashmap

import "github.com/katoozi/avran_cloud_code_challenge_wallet_service/internal/models"

func (hm *hashMap) Balance(cellphone string) (int64, error) {
	item, ok := hm.users.Load(cellphone)
	if !ok {
		item = new(models.Wallet)
		hm.users.Store(cellphone, item)
	}

	wallet := item.(*models.Wallet)

	return wallet.CurrentBalance, nil
}
