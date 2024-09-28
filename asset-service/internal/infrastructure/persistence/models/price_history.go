package models

import (
	"time"

	"goldvault/asset-service/internal/core/domain/entity"
)

type PriceHistory struct {
	ID        int64
	AssetType string
	BuyPrice  float64
	SellPrice float64
	CreatedAt time.Time
}

const (
	PricesHistoryTableName  = "price_histories"
	PriceHistoryColumns     = "id, asset_type, buy_price, sell_price, created_at"
	PriceHistoryColumnsNoID = "asset_type, buy_price, sell_price, created_at"
)

func (p *PriceHistory) Scan(scanner Scanner) error {
	return scanner.Scan(&p.ID, &p.AssetType, &p.BuyPrice, &p.SellPrice, &p.CreatedAt)
}

func (p *PriceHistory) ToEntity() *entity.PriceHistory {
	return &entity.PriceHistory{
		ID:        p.ID,
		AssetType: entity.AssetType(p.AssetType),
		Prices: entity.PriceDetails{
			BuyPrice:  p.BuyPrice,
			SellPrice: p.SellPrice,
		},
		CreatedAt: p.CreatedAt,
	}
}

func ToPriceHistoryDB(priceHistory *entity.PriceHistory) *PriceHistory {
	return &PriceHistory{
		ID:        priceHistory.ID,
		AssetType: priceHistory.AssetType.String(),
		BuyPrice:  priceHistory.Prices.BuyPrice,
		SellPrice: priceHistory.Prices.SellPrice,
		CreatedAt: priceHistory.CreatedAt,
	}
}
