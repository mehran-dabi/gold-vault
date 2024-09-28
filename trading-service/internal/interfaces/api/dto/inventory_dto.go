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
