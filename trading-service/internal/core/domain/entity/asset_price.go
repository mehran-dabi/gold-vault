package entity

import (
	"fmt"
	"time"
)

type AssetType string

const (
	AssetTypeGold AssetType = "gold"
)

func (a AssetType) Validate() error {
	switch a {
	case AssetTypeGold:
		return nil
	default:
		return fmt.Errorf("invalid asset type")
	}
}

func (a AssetType) String() string {
	return string(a)
}

type AssetPrice struct {
	ID        int64
	AssetType AssetType
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *AssetPrice) Validate() error {
	if a.Price < 0 {
		return fmt.Errorf("price must be greater than or equal to 0")
	}
	return nil
}

type PriceDetails struct {
	BuyPrice  float64
	SellPrice float64
}
