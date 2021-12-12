package models

import "sync"

// Wallet is the structure of users wallet
type Wallet struct {
	CurrentBalance int64
	Transactions   sync.Map
}
