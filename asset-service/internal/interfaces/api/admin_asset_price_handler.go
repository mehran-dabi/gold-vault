package api

import (
	"net/http"

	"goldvault/asset-service/internal/core/application/services"
	"goldvault/asset-service/internal/core/domain/entity"
	"goldvault/asset-service/internal/interfaces/api/dto"
	"goldvault/asset-service/pkg/serr"

	"github.com/gin-gonic/gin"
)

type AdminAssetPriceHandler struct {
	assetPriceService *services.AssetPriceService
}

func NewAdminAssetPriceHandler(assetPriceService *services.AssetPriceService) *AdminAssetPriceHandler {
	return &AdminAssetPriceHandler{
		assetPriceService: assetPriceService,
	}
}

// UpsertPrice godoc
// @Summary Upsert asset price
// @Description Adds or updates the price of a specific asset for admins.
// @Tags Admin Asset Prices
// @Accept json
// @Produce json
// @Param request body dto.UpsertAssetPrice true "Upsert asset price request"
// @Success 200 {object} map[string]string "Price upserted successfully"
// @Failure 400 {object} Error "Invalid input"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/asset-prices [post]
func (a *AdminAssetPriceHandler) UpsertPrice(c *gin.Context) {
	var req dto.UpsertAssetPrice
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, serr.ValidationErr("AdminAssetPriceHandler.UpsertPrice", "invalid input", serr.ErrInvalidInput))
		return
	}

	err := req.Validate()
	if err != nil {
		handleError(c, serr.ValidationErr("AdminAssetPriceHandler.UpsertPrice", err.Error(), serr.ErrInvalidInput))
		return
	}

	err = a.assetPriceService.UpsertPrice(c.Request.Context(), req.AssetType, &entity.PriceDetails{BuyPrice: req.BuyPrice, SellPrice: req.SellPrice})
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "price upserted"})
}

// DeleteAssetPrice godoc
// @Summary Delete asset price
// @Description Deletes the price of a specific asset for admins.
// @Tags Admin Asset Prices
// @Accept json
// @Produce json
// @Param assetType path string true "Asset Type"
// @Success 200 {object} map[string]string "Price deleted successfully"
// @Failure 400 {object} Error "Invalid asset ID"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 404 {object} Error "Asset price not found"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/asset-prices/{assetType} [delete]
func (a *AdminAssetPriceHandler) DeleteAssetPrice(c *gin.Context) {
	assetType := c.Param("asset_type")
	if assetType == "" {
		handleError(c, serr.ValidationErr("AdminAssetPriceHandler.DeleteAssetPrice", "asset type is required", serr.ErrInvalidInput))
		return
	}
	err := a.assetPriceService.DeleteAssetPrice(c.Request.Context(), assetType)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "price deleted"})
}
