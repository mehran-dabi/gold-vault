package ports

import "context"

type (
	PriceCachePorts interface {
		SavePrice(ctx context.Context, assetType string, price float64) error
		GetPrice(ctx context.Context, assetType string) (float64, error)
		RemovePrice(ctx context.Context, assetType string) error
	}
)
