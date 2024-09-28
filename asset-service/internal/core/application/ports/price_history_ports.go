package ports

import (
	"context"

	"goldvault/asset-service/internal/core/domain/entity"
)

type (
	PriceHistoryPersistence interface {
		GetHistoryByAssetType(ctx context.Context, assetType string, limit, offset int64) ([]*entity.PriceHistory, error)
		InsertHistory(ctx context.Context, priceHistory *entity.PriceHistory) error
	}

	PriceHistoryDomainService interface {
		GetAssetPriceHistory(ctx context.Context, assetType string, limit, offset int64) ([]*entity.PriceHistory, error)
		AddPriceHistory(ctx context.Context, assetType string, prices *entity.PriceDetails) error
	}
)
