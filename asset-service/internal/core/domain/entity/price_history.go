package entity

import (
	"fmt"
	"time"
)

type PriceHistory struct {
	ID        int64
	AssetType AssetType
	Prices    PriceDetails
	CreatedAt time.Time
}

func (p *PriceHistory) Validate() error {
	if p.AssetType == "" {
		return fmt.Errorf("asset type must not be empty")
	}

	if p.Prices.BuyPrice < 0 {
		return fmt.Errorf("buy price must be greater than or equal to 0")
	}

	if p.Prices.SellPrice < 0 {
		return fmt.Errorf("sell price must be greater than or equal to 0")
	}

	return nil
}
