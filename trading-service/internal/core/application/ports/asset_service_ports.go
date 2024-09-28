package ports

import (
	"context"

	"goldvault/trading-service/internal/core/domain/entity"
)

type (
	AssetServiceClient interface {
		GetAssetPrice(ctx context.Context, assetTypes []string) (map[string]entity.PriceDetails, error)
	}
)
