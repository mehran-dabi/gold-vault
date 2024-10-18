package dto

type (
	BuyAssetRequest struct {
		AssetType string  `json:"asset_type" validate:"required"`
		Amount    float64 `json:"amount" validate:"required,gt=0"`
		Price     float64 `json:"price" validate:"required,gt=0"`
	}

	SellAssetRequest struct {
		AssetType string  `json:"asset_type" validate:"required"`
		Amount    float64 `json:"amount" validate:"required,gt=0"`
		Price     float64 `json:"price" validate:"required,gt=0"`
	}

	CreateInventoryAdminRequest struct {
		AssetType string  `json:"asset_type" validate:"required"`
		Quantity  float64 `json:"quantity" validate:"required,gt=0"`
	}

	SetGlobalLimitsRequest struct {
		MinBuy         float64 `json:"min_buy" validate:"required,gt=0"`
		MaxBuy         float64 `json:"max_buy" validate:"required,gt=0"`
		MinSell        float64 `json:"min_sell" validate:"required,gt=0"`
		MaxSell        float64 `json:"max_sell" validate:"required,gt=0"`
		DailyBuyLimit  float64 `json:"daily_buy_limit" validate:"required,gt=0"`
		DailySellLimit float64 `json:"daily_sell_limit" validate:"required,gt=0"`
	}
)

func (b *BuyAssetRequest) Validate() error {
	return GetValidator().Struct(b)
}

func (s *SellAssetRequest) Validate() error {
	return GetValidator().Struct(s)
}

func (c *CreateInventoryAdminRequest) Validate() error {
	return GetValidator().Struct(c)
}
