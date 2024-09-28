package dto

type (
	GetAssetPriceHistoryRequest struct {
		AssetType string `json:"asset_type" validate:"required"`
		Limit     int64  `json:"limit" validate:"required"`
		Offset    int64  `json:"offset" validate:"required"`
	}
)

func (g *GetAssetPriceHistoryRequest) Validate() error {
	return GetValidator().Struct(g)
}
