package hashmap

import (
	"errors"
	"testing"
	"time"
)

var hm = &hashMap{}

func Test_hashMap_Transaction(t *testing.T) {
	type args struct {
		cellphone   string
		amount      int64
		timestamp   time.Time
		description string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{name: "#1", args: args{cellphone: "09017806181", amount: 1000, timestamp: time.Now(), description: "test trx"}, wantErr: nil},
		{name: "#2", args: args{cellphone: "09017806181", amount: 23000, timestamp: time.Now(), description: "test trx"}, wantErr: nil},
		{name: "#3", args: args{cellphone: "09017806181", amount: 47500, timestamp: time.Now(), description: "test trx"}, wantErr: nil},
		{name: "#4", args: args{cellphone: "09017806181", amount: -10000, timestamp: time.Now(), description: "test trx"}, wantErr: nil},
		{name: "#5", args: args{cellphone: "09017806181", amount: 112000, timestamp: time.Now(), description: "test trx"}, wantErr: nil},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, _, err := hm.Transaction(tt.args.cellphone, tt.args.amount, tt.args.timestamp, tt.args.description)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("hashMap.Transaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

	t.Cleanup(func() {})
}

func Test_hashMap_BalanceAfterConcurrentTransactions(t *testing.T) {
	t.Run("#6. test balance after concurrent increments", func(t *testing.T) {
		balance, err := hm.Balance("09017806181")
		if !errors.Is(err, nil) {
			t.Errorf("hashMap.Balance() error = %v, wantErr %v", err, nil)
		}

		var wantedBalance int64 = 1000 + 23000 + 47500 - 10000 + 112000
		if balance != wantedBalance {
			t.Errorf("hashMap.Balance() balance = %v, wantBalance %v", balance, wantedBalance)
		}
	})
}
