package hashmap

import (
	"sync/atomic"
	"time"

	discountModels "github.com/katoozi/avran_cloud_code_challenge_discount_service/internal/discount/models"
)

func (hm *hashMap) RedeemVoucher(cellphone, voucherCode string) (uint64, error) {
	item, ok := hm.vouchers.Load(voucherCode)
	if !ok {
		return 0, ErrVoucherNotFound
	}

	voucher := item.(*storageItem)

	if _, ok := voucher.redemptionRecords.Load(cellphone); ok {
		return 0, ErrDoubleRedemption
	}

	// * for the sake of performance, we are using atomic counters instead of mutexes

	newTotalMaxUseTime := atomic.AddUint64(&voucher.totalMaxUsedTime, 1)

	if newTotalMaxUseTime > voucher.maxUseTime {
		// * we can reset the counter to the previous value.
		// * but we can use this value as attempts to redeem the voucher
		// ! here is the solution to reset the value
		// value := newTotalMaxUseTime % voucher.maxUseTime
		// if value > 0 {
		// 	atomic.StoreUint32(&voucher.totalMaxUsedTime, voucher.maxUseTime+1) // Reset the counter to max value
		// }

		return 0, ErrVoucherIsOver
	}

	voucher.redemptionRecords.Store(cellphone, time.Now())

	return voucher.discountPrice, nil
}

func (hm *hashMap) VoucherRedemptionRecords(voucherCode string) ([]discountModels.RedemptionRecord, error) {
	item, ok := hm.vouchers.Load(voucherCode)
	if !ok {
		return nil, ErrVoucherNotFound
	}
	voucher := item.(*storageItem)

	result := make([]discountModels.RedemptionRecord, 0)
	voucher.redemptionRecords.Range(func(key, value interface{}) bool {
		result = append(result, discountModels.RedemptionRecord{Cellphone: key.(string), Timestamp: value.(time.Time)})
		return true
	})

	return result, nil
}
