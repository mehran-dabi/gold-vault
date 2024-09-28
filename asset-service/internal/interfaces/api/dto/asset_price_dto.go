package dto

type (
	UpsertAssetPrice struct {
		AssetType string  `json:"asset_type" validate:"required"`
		BuyPrice  float64 `json:"buy_price" validate:"required"`
		SellPrice float64 `json:"sell_price" validate:"required"`
	}
)

func (u *UpsertAssetPrice) Validate() error {
	return GetValidator().Struct(u)
}
