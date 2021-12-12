package hashmap

import (
	discountModels "github.com/katoozi/avran_cloud_code_challenge_discount_service/internal/discount/models"
)

func (hm *hashMap) SaveVoucher(voucher discountModels.Voucher) error {
	if _, ok := hm.vouchers.Load(voucher.Code); ok {
		return ErrDuplicateVoucherCode
	}
	hm.vouchers.Store(voucher.Code, newStorageItem(voucher.CreatedAt, voucher.DiscountPrice, voucher.MaxUseTime, 0))
	return nil
}
