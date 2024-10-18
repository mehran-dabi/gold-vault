package ports

import (
	"context"

	"goldvault/asset-service/internal/core/domain/entity"
)

type (
	AssetPricePersistence interface {
		GetPrice(ctx context.Context, assetType string) (*entity.PriceDetails, error)
		UpsertPrice(ctx context.Context, assetType string, prices *entity.PriceDetails) error
		DeleteAssetPrice(ctx context.Context, assetType string) error
		GetAllAssetPrices(ctx context.Context) (map[string]*entity.PriceDetails, error)
		UpdateAssetPriceByStep(ctx context.Context, step float64, assetType string) error
	}

	AssetPriceDomainService interface {
		GetLatestPrice(ctx context.Context, assetType string) (*entity.PriceDetails, error)
		UpsertPrice(ctx context.Context, assetType string, prices *entity.PriceDetails) error
		DeleteAssetPrice(ctx context.Context, assetType string) error
		GetAllAssetPrices(ctx context.Context) (map[string]*entity.PriceDetails, error)
		UpdateAssetPriceByStep(ctx context.Context, step float64, assetType string) error
	}
)
