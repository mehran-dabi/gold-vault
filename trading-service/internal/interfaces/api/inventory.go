package api

import (
	"net/http"

	"goldvault/trading-service/internal/core/application/services"
	"goldvault/trading-service/internal/interfaces/api/dto"
	"goldvault/trading-service/pkg/serr"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	inventoryService *services.InventoryService
}

func NewInventoryHandler(inventoryService *services.InventoryService) *InventoryHandler {
	return &InventoryHandler{
		inventoryService: inventoryService,
	}
}

// BuyAsset godoc
// @Summary Buy an asset
// @Description Allows a user to buy a specific asset by providing asset type, amount, and price.
// @Tags Inventory
// @Accept json
// @Produce json
// @Param request body dto.BuyAssetRequest true "Buy asset request"
// @Success 200 {object} map[string]string "Asset bought successfully"
// @Failure 400 {object} Error "Invalid input"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /trades/buy [post]
func (i *InventoryHandler) BuyAsset(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		handleError(c, serr.ValidationErr("InventoryHandler.BuyAsset", "invalid user ID", serr.ErrInvalidInput))
		return
	}

	var req dto.BuyAssetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, serr.ValidationErr("InventoryHandler.BuyAsset", err.Error(), serr.ErrInvalidInput))
		return
	}

	err := i.inventoryService.BuyAsset(c.Request.Context(), userID, req.AssetType, req.Amount, req.Price)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "asset bought successfully"})
}

// SellAsset godoc
// @Summary Sell an asset
// @Description Allows a user to sell a specific asset by providing asset type, amount, and price.
// @Tags Inventory
// @Accept json
// @Produce json
// @Param request body dto.SellAssetRequest true "Sell asset request"
// @Success 200 {object} map[string]string "Asset sold successfully"
// @Failure 400 {object} Error "Invalid input"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /trades/sell [post]
func (i *InventoryHandler) SellAsset(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		handleError(c, serr.ValidationErr("InventoryHandler.SellAsset", "invalid user ID", serr.ErrInvalidInput))
		return
	}

	var req dto.SellAssetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, serr.ValidationErr("InventoryHandler.SellAsset", err.Error(), serr.ErrInvalidInput))
		return
	}

	err := i.inventoryService.SellAsset(c.Request.Context(), userID, req.AssetType, req.Amount, req.Price)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "asset sold successfully"})
}
