package ports

import (
	"context"

	"goldvault/wallet-service/internal/core/domain/entity"
)

type (
	AssetServiceClientPorts interface {
		GetAssetPrice(ctx context.Context, assetTypes []string) (map[string]*entity.PriceDetails, error)
	}
)
