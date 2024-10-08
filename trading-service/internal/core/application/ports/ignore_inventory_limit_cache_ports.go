package ports

import "context"

type (
	IgnoreInventoryLimitCache interface {
		Set(ctx context.Context) error
		Get(ctx context.Context) (bool, error)
	}
)
