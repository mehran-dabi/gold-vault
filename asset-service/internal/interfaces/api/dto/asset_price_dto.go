package dto

type (
	UpsertAssetPrice struct {
		AssetType string  `json:"asset_type" validate:"required"`
		BuyPrice  float64 `json:"buy_price" validate:"required"`
		SellPrice float64 `json:"sell_price" validate:"required"`
	}

	SetPriceChangeStep struct {
		Step float64 `json:"step" validate:"required"`
	}
)

func (u *UpsertAssetPrice) Validate() error {
	return GetValidator().Struct(u)
}

func (s *SetPriceChangeStep) Validate() error {
	return GetValidator().Struct(s)
}
