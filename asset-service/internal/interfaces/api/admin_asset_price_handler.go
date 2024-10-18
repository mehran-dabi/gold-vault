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

// UpdateAssetPriceByStep godoc
// @Summary Update asset price by step
// @Description Updates the price of an asset by a predefined step for the specified asset type.
// @Tags Admin Asset Price
// @Produce json
// @Param assetType path string true "Asset type"
// @Success 200 {object} map[string]string "Message indicating the price was updated"
// @Failure 400 {object} Error "Invalid asset type or missing parameters"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/asset-prices/{asset_type}/adjust-by-step [put]
func (a *AdminAssetPriceHandler) UpdateAssetPriceByStep(c *gin.Context) {
	assetType := c.Param("assetType")
	if assetType == "" {
		handleError(c, serr.ValidationErr("AdminAssetPriceHandler.UpdateAssetPriceByStep", "asset type is required", serr.ErrInvalidInput))
		return
	}

	err := a.assetPriceService.UpdateAssetPriceByStep(c.Request.Context(), assetType)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "price updated"})
}

// SetPriceChangeStep godoc
// @Summary Set price change step
// @Description Sets the step value for price changes.
// @Tags Admin Asset Price
// @Accept json
// @Produce json
// @Param body body dto.SetPriceChangeStep true "Set price change step"
// @Success 200 {object} map[string]string "Message indicating the price change step was set"
// @Failure 400 {object} Error "Invalid input"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/asset-prices/step [post]
func (a *AdminAssetPriceHandler) SetPriceChangeStep(c *gin.Context) {
	var req dto.SetPriceChangeStep
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, serr.ValidationErr("AdminAssetPriceHandler.SetPriceChangeStep", "invalid input", serr.ErrInvalidInput))
		return
	}

	err := req.Validate()
	if err != nil {
		handleError(c, serr.ValidationErr("AdminAssetPriceHandler.SetPriceChangeStep", err.Error(), serr.ErrInvalidInput))
		return
	}

	err = a.assetPriceService.SetPriceChangeStep(c.Request.Context(), req.Step)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "price change step set"})
}

// GetPriceChangeStep godoc
// @Summary Get price change step
// @Description Retrieves the current step value used for price changes.
// @Tags Admin Asset Price
// @Produce json
// @Success 200 {object} map[string]float64 "The current price change step"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/asset-prices/step [get]
func (a *AdminAssetPriceHandler) GetPriceChangeStep(c *gin.Context) {
	step, err := a.assetPriceService.GetPriceChangeStep(c.Request.Context())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"step": step})
}
