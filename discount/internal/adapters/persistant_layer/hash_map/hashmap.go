package hashmap

import (
	"sync"

	"github.com/katoozi/avran_cloud_code_challenge_discount_service/internal/ports"
)

// hashMap persistant layer will implement the PersistantLayer port
type hashMap struct {
	// * sync.map is more performant here because we are mostly reading vouchers instead of writing vouchers
	// * we can also use libraries like badger for an embedded cache. https://github.com/dgraph-io/badger
	vouchers sync.Map
}

// New is the HashMap persistant layer factory method
func New() ports.PersistantLayer {
	return &hashMap{
		vouchers: sync.Map{},
	}
}
