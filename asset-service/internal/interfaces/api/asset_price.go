package api

import (
	"net/http"

	"goldvault/asset-service/internal/core/application/services"
	"goldvault/asset-service/pkg/serr"

	"github.com/gin-gonic/gin"
)

type AssetPriceHandler struct {
	assetPriceService *services.AssetPriceService
}

func NewAssetPriceHandler(assetPriceService *services.AssetPriceService) *AssetPriceHandler {
	return &AssetPriceHandler{
		assetPriceService: assetPriceService,
	}
}

// GetPrice godoc
// @Summary Get latest asset price
// @Description Retrieves the latest price for a specific asset type.
// @Tags Asset Prices
// @Accept json
// @Produce json
// @Param assetType path string true "Asset type (e.g., gold, silver)"
// @Success 200 {object} map[string]float64 "Latest price retrieved successfully"
// @Failure 400 {object} Error "Asset type is required or invalid input"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 404 {object} Error "Asset not found"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /asset-prices/{assetType} [get]
func (a *AssetPriceHandler) GetPrice(c *gin.Context) {
	assetType := c.Param("assetType")
	if assetType == "" {
		handleError(c, serr.ValidationErr("AssetPriceHandler.GetPrice",
			"asset type is required", serr.ErrInvalidInput))
		return
	}

	price, err := a.assetPriceService.GetLatestPrice(c.Request.Context(), assetType)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"price": price})
}

// GetAllAssetPrices godoc
// @Summary Get all asset prices
// @Description Retrieves the latest prices for all available assets.
// @Tags Asset Prices
// @Accept json
// @Produce json
// @Success 200 {object} map[string]entity.PriceDetails "Latest prices for all assets retrieved successfully"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /asset-prices [get]
func (a *AssetPriceHandler) GetAllAssetPrices(c *gin.Context) {
	prices, err := a.assetPriceService.GetAllAssetPrices(c.Request.Context())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, prices)
}
