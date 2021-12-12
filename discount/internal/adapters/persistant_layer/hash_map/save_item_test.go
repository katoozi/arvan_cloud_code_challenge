package hashmap

import (
	"errors"
	"testing"

	discountModels "github.com/katoozi/avran_cloud_code_challenge_discount_service/internal/discount/models"
)

func Test_hashMap_SaveVoucher(t *testing.T) {
	type args struct {
		voucher discountModels.Voucher
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{name: "#1", args: args{voucher: discountModels.Voucher{Code: "test_1"}}, wantErr: nil},
		{name: "#2", args: args{voucher: discountModels.Voucher{Code: "test_1"}}, wantErr: ErrDuplicateVoucherCode},
	}
	hm := &hashMap{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := hm.SaveVoucher(tt.args.voucher)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("hashMap.SaveVoucher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
