package api

import (
	"net/http"

	"goldvault/trading-service/internal/core/application/services"
	"goldvault/trading-service/internal/interfaces/api/dto"
	"goldvault/trading-service/pkg/serr"

	"github.com/gin-gonic/gin"
)

type InventoryAdminHandler struct {
	inventoryAdminService *services.InventoryService
}

func NewInventoryAdminHandler(inventoryAdminService *services.InventoryService) *InventoryAdminHandler {
	return &InventoryAdminHandler{
		inventoryAdminService: inventoryAdminService,
	}
}

// CreateInventory godoc
// @Summary Create new inventory
// @Description Allows an admin to create a new inventory by specifying the asset type and quantity.
// @Tags Admin Inventory
// @Accept json
// @Produce json
// @Param request body dto.CreateInventoryAdminRequest true "Create inventory request"
// @Success 200 {object} map[string]int "Inventory created successfully with ID"
// @Failure 400 {object} Error "Invalid input"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/inventory [post]
func (i *InventoryAdminHandler) CreateInventory(c *gin.Context) {
	var req dto.CreateInventoryAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, serr.ValidationErr("InventoryHandler.CreateInventory", err.Error(), serr.ErrInvalidInput))
		return
	}

	// create inventory entity
	id, err := i.inventoryAdminService.CreateInventory(c.Request.Context(), req.AssetType, req.Quantity)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// DeleteInventory godoc
// @Summary Delete an inventory
// @Description Allows an admin to delete an inventory by asset type.
// @Tags Admin Inventory
// @Accept json
// @Produce json
// @Param assetType path string true "Asset type to delete from inventory"
// @Success 200 {object} map[string]string "Inventory deleted successfully"
// @Failure 400 {object} Error "Invalid asset type"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 404 {object} Error "Asset not found"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/inventory/{assetType} [delete]
func (i *InventoryAdminHandler) DeleteInventory(c *gin.Context) {
	assetType := c.Param("assetType")
	if assetType == "" {
		handleError(c, serr.ValidationErr("InventoryHandler.DeleteInventory", "asset type is required", serr.ErrInvalidInput))
		return
	}

	err := i.inventoryAdminService.DeleteInventory(c.Request.Context(), assetType)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "inventory deleted successfully"})
}

// GetInventory godoc
// @Summary Get all inventory
// @Description Retrieves the current inventory for all asset types.
// @Tags Admin Inventory
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]entity.Inventory "Inventory retrieved successfully"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/inventory [get]
func (i *InventoryAdminHandler) GetInventory(c *gin.Context) {
	inventory, err := i.inventoryAdminService.GetInventory(c.Request.Context())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"inventory": inventory})
}
