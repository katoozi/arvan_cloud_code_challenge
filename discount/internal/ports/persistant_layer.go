package ports

import (
	discountModels "github.com/katoozi/avran_cloud_code_challenge_discount_service/internal/discount/models"
)

// PersistantLayer will wrap the vouchers storage behavior
type PersistantLayer interface {
	SaveVoucher(discountModels.Voucher) error
	RedeemVoucher(cellphone, voucherCode string) (uint64, error)
	VoucherRedemptionRecords(voucherCode string) ([]discountModels.RedemptionRecord, error)
}
