package hashmap

import (
	"errors"
	"sync"
	"testing"

	"github.com/katoozi/avran_cloud_code_challenge_wallet_service/internal/models"
)

func Test_hashMap_Balance(t *testing.T) {
	usersSyncMap := sync.Map{}
	usersSyncMap.Store("09017806182", &models.Wallet{CurrentBalance: 10000})
	hm := &hashMap{
		users: usersSyncMap,
	}

	type args struct {
		cellphone string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr error
	}{
		{name: "#1", args: args{cellphone: "09017806181"}, want: 0, wantErr: nil},
		{name: "#2", args: args{cellphone: "09017806182"}, want: 10000, wantErr: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hm.Balance(tt.args.cellphone)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("hashMap.Balance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hashMap.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}
