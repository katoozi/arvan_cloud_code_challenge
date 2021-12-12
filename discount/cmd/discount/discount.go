package main

import (
	"log"
	"net/http"
	"time"

	"github.com/katoozi/avran_cloud_code_challenge_discount_service/internal/adapters/notifier/wallet"
	hashmap "github.com/katoozi/avran_cloud_code_challenge_discount_service/internal/adapters/persistant_layer/hash_map"
	discountService "github.com/katoozi/avran_cloud_code_challenge_discount_service/internal/discount"
)

func main() {
	// TODO: use environment variables or a config file for configurations

	// ! we can implement a new storage like: redis, mysql, postgresql, etcd, etc...
	storage := hashmap.New()

	// * i am going to call the wallet service through the rest api
	// ! we can implement a new notifier like mysql, postgresql, etcd, redis queues, nats, rabbitmq, kafka, etc...
	notifier := wallet.New("http://localhost:8081", time.Second*1)

	// TODO: handle input data validation and http methods. consider using a library or framework.
	http.HandleFunc("/redeem", discountService.RedeemHandler(storage, notifier))
	http.HandleFunc("/voucher", discountService.CreateVoucherHandler(storage, notifier))
	http.HandleFunc("/voucher/redemptions", discountService.VoucherRedemptionsHandler(storage, notifier))

	// TODO: impelement gracefull shutdown.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
