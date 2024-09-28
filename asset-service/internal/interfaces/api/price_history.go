package api

import (
	"net/http"
	"strconv"

	"goldvault/asset-service/internal/core/application/services"

	"github.com/gin-gonic/gin"
)

type PriceHistoryHandler struct {
	priceHistoryService *services.PriceHistoryService
}

func NewPriceHistoryHandler(priceHistoryService *services.PriceHistoryService) *PriceHistoryHandler {
	return &PriceHistoryHandler{
		priceHistoryService: priceHistoryService,
	}
}

// GetAssetPriceHistory godoc
// @Summary Get asset price history
// @Description Retrieves the price history for a specific asset type with optional pagination.
// @Tags Asset Price History
// @Accept json
// @Produce json
// @Param assetType path string true "Asset type (e.g., gold, silver)"
// @Param limit query int false "Limit the number of results (default is 10)"
// @Param offset query int false "Offset for pagination (default is 0)"
// @Success 200 {object} map[string][]entity.PriceHistory "Price history retrieved successfully"
// @Failure 400 {object} Error "Invalid limit or offset"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 404 {object} Error "Asset price history not found"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /price-history/{assetType} [get]
func (p *PriceHistoryHandler) GetAssetPriceHistory(c *gin.Context) {
	assetType := c.Param("assetType")

	limit := 10
	offset := 0

	if l := c.Query("limit"); l != "" {
		parsedLimit, err := strconv.Atoi(l)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if o := c.Query("offset"); o != "" {
		parsedOffset, err := strconv.Atoi(o)
		if err == nil && parsedOffset > 0 {
			offset = parsedOffset
		}
	}

	priceHistory, err := p.priceHistoryService.GetAssetPriceHistory(c.Request.Context(), assetType, int64(limit), int64(offset))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"price_history": priceHistory})
}
