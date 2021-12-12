package ports

import (
	"context"
	"time"
)

// Notifier will wrap the transaction system behavior
type Notifier interface {
	Transaction(ctx context.Context, cellphone string, amount uint64, description string, timestamp time.Time) error
}
