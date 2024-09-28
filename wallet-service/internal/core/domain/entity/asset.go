package entity

import "time"

type Asset struct {
	ID         int64     `json:"id"`
	WalletID   int64     `json:"wallet_id"`
	Type       string    `json:"type"`        // Type of asset, e.g., 'gold', 'USD'
	Balance    float64   `json:"balance"`     // Asset balance
	TotalPrice float64   `json:"total_price"` // Total price of the asset
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type PriceDetails struct {
	BuyPrice  float64
	SellPrice float64
}

const (
	AssetTypeIRR = "IRR"
)
