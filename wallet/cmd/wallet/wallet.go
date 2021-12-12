package main

import (
	"log"
	"net/http"

	"github.com/katoozi/avran_cloud_code_challenge_wallet_service/internal/adapters/persistant_layer/hashmap"
	walletService "github.com/katoozi/avran_cloud_code_challenge_wallet_service/internal/wallet"
)

func main() {
	storage := hashmap.New()

	http.HandleFunc("/trx", walletService.TransactionHandler(storage))
	http.HandleFunc("/balance", walletService.BalanceHandler(storage))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
