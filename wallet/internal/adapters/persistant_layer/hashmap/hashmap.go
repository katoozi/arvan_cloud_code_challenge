package hashmap

import (
	"sync"

	"github.com/katoozi/avran_cloud_code_challenge_wallet_service/internal/ports"
)

type hashMap struct {
	users          sync.Map
	transactionIDs sync.Map
}

// New is the hashmap persistant layer factory method
func New() ports.PersistantLayer {
	return new(hashMap)
}
