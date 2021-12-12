package ports

import (
	"errors"
	"time"

	"github.com/katoozi/avran_cloud_code_challenge_wallet_service/internal/models"
)

// PersistantLayer will wrap the storage behavior
type PersistantLayer interface {
	Transaction(cellphone string, amount int64, timestamp time.Time, description string) (models.Transaction, int64, error)
	Balance(cellphone string) (int64, error)
}

var (
	// ErrPersistantLayerUserNotFound will returned when user is not exists.
	ErrPersistantLayerUserNotFound = errors.New("user not found")
)
