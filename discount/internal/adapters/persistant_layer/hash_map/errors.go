package hashmap

import "errors"

var (
	// ErrDuplicateVoucherCode will be returned if you want to save a duplicate voucher code
	ErrDuplicateVoucherCode error = errors.New("duplicate voucher code")
	// ErrDoubleRedemption will be returned if a cellphone want's to redeem a voucher twice.
	ErrDoubleRedemption error = errors.New("double redemption")
	// ErrVoucherNotFound will be returned if you want to redeem a voucher that is not available
	ErrVoucherNotFound error = errors.New("voucher code not found")
	// ErrVoucherIsOver will be returned if you want to redeem a voucher that is not available
	ErrVoucherIsOver error = errors.New("voucher code is over")
)
