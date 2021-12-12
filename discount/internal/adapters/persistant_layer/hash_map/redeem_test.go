package hashmap

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func Test_hashMap_RedeemVoucher(t *testing.T) {
	syncMap := sync.Map{}
	type args struct {
		cellphone   string
		voucherCode string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr error

		applyChanges  func()
		revertChanges func()
	}{
		{name: "#1", args: args{cellphone: "09017806181", voucherCode: "voucher_1"}, want: 0, wantErr: ErrVoucherNotFound, applyChanges: func() {}, revertChanges: func() {}},

		{name: "#2", args: args{cellphone: "09017806182", voucherCode: "voucher_2"}, want: 100000, wantErr: nil, applyChanges: func() {
			syncMap.Store("voucher_2", newStorageItem(time.Now(), 100000, 1, 0))
		}, revertChanges: func() {
			syncMap.Delete("voucher_2")
		}},

		{name: "#3", args: args{cellphone: "09017806183", voucherCode: "voucher_3"}, want: 0, wantErr: ErrDoubleRedemption, applyChanges: func() {
			voucher := newStorageItem(time.Now(), 100000, 1, 0)
			voucher.redemptionRecords.Store("09017806183", time.Now())
			syncMap.Store("voucher_3", voucher)
		}, revertChanges: func() {
			syncMap.Delete("voucher_3")
		}},

		{name: "#4", args: args{cellphone: "09017806184", voucherCode: "voucher_4"}, want: 0, wantErr: ErrVoucherIsOver, applyChanges: func() {
			syncMap.Store("voucher_4", newStorageItem(time.Now(), 100000, 0, 0))
		}, revertChanges: func() {
			syncMap.Delete("voucher_4")
		}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // also check the concurrent safety

			tt.applyChanges()
			hm := &hashMap{
				vouchers: syncMap,
			}

			got, err := hm.RedeemVoucher(tt.args.cellphone, tt.args.voucherCode)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("hashMap.RedeemVoucher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hashMap.RedeemVoucher() = %v, want %v", got, tt.want)
			}
			tt.revertChanges()
		})
	}

	t.Cleanup(func() {})
}
