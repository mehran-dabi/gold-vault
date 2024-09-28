package models

import (
	"time"

	"goldvault/asset-service/internal/core/domain/entity"
)

type AssetPrice struct {
	ID        int64
	AssetType string
	BuyPrice  float64
	SellPrice float64
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	AssetPriceTableName   = "asset_prices"
	AssetPriceColumns     = "id, asset_type, buy_price, sell_price, created_at, updated_at"
	AssetPriceColumnsNoID = "asset_type, buy_price, sell_price, created_at, updated_at"
)

func (a *AssetPrice) Scan(scanner Scanner) error {
	return scanner.Scan(
		&a.ID,
		&a.AssetType,
		&a.BuyPrice,
		&a.SellPrice,
		&a.CreatedAt,
		&a.UpdatedAt,
	)
}

func ToAssetPriceDB(entity *entity.AssetPrice) *AssetPrice {
	return &AssetPrice{
		ID:        entity.ID,
		AssetType: entity.AssetType.String(),
		BuyPrice:  entity.Prices.BuyPrice,
		SellPrice: entity.Prices.SellPrice,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func (a *AssetPrice) ToEntity() *entity.AssetPrice {
	return &entity.AssetPrice{
		ID:        a.ID,
		AssetType: entity.AssetType(a.AssetType),
		Prices: entity.PriceDetails{
			BuyPrice:  a.BuyPrice,
			SellPrice: a.SellPrice,
		},
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}
