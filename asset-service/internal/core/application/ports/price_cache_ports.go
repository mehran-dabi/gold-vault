package ports

import (
	"context"

	"goldvault/asset-service/internal/core/domain/entity"
)

type (
	PriceCachePorts interface {
		SavePrice(ctx context.Context, assetType string, prices *entity.PriceDetails) error
		GetPrice(ctx context.Context, assetType string) (*entity.PriceDetails, error)
		RemovePrice(ctx context.Context, assetType string) error
		GetAllAssetPrices(ctx context.Context) (map[string]*entity.PriceDetails, error)
		GetPriceChangeStep(ctx context.Context) (float64, error)
		SetPriceChangeStep(ctx context.Context, step float64) error
	}
)
