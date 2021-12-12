package wallet

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/katoozi/avran_cloud_code_challenge_wallet_service/internal/ports"
)

// Service is the wallet service structure
type Service struct {
	wallets         sync.Map
	persistantLayer ports.PersistantLayer
}

// New is the wallet service factory method
func New(persistantLayer ports.PersistantLayer) *Service {
	return &Service{
		persistantLayer: persistantLayer,
	}
}

// TransactionHandler will inject dependencies and return the redeem http handler
func TransactionHandler(persistantLayer ports.PersistantLayer) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		cellphone := req.URL.Query().Get("cellphone")
		description := req.URL.Query().Get("description")
		trxType := req.URL.Query().Get("type")
		amount, _ := strconv.ParseInt(req.URL.Query().Get("amount"), 10, 64)
		timestamp, _ := time.Parse(time.RFC3339, req.URL.Query().Get("timestamp"))

		// there two types: 1 for increasing the balance and 2 for decreasing it.
		if trxType == "2" {
			amount = -1 * amount
		}

		trx, balance, err := persistantLayer.Transaction(cellphone, amount, timestamp, description)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]interface{}{"trx": trx, "balance": balance})
	}
}

// BalanceHandler will inject dependencies and return the get balance http handler
func BalanceHandler(persistantLayer ports.PersistantLayer) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		cellphone := req.URL.Query().Get("cellphone")

		balance, err := persistantLayer.Balance(cellphone)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write([]byte(strconv.FormatInt(balance, 10)))
	}
}
