package api

import (
	"net/http"
	"strconv"

	"goldvault/wallet-service/internal/core/application/services"
	"goldvault/wallet-service/pkg/serr"

	"github.com/gin-gonic/gin"
)

type AdminWalletHandler struct {
	walletService *services.WalletService
}

func NewAdminWalletHandler(walletService *services.WalletService) *AdminWalletHandler {
	return &AdminWalletHandler{
		walletService: walletService,
	}
}

// GetWallets godoc
// @Summary Get all wallets with pagination
// @Description Retrieves a list of wallets with pagination for admins.
// @Tags Admin Wallets
// @Accept json
// @Produce json
// @Param limit query int false "Limit the number of results"
// @Param offset query int false "Offset for pagination"
// @Success 200 {array} entity.Wallet "List of wallets retrieved successfully"
// @Failure 400 {object} Error "Invalid pagination parameters"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/wallets [get]
func (a *AdminWalletHandler) GetWallets(c *gin.Context) {
	limit := 10
	offset := 0

	if l := c.Query("limit"); l != "" {
		parsedLimit, err := strconv.Atoi(l)
		if err != nil {
			handleError(c, serr.ValidationErr("handler.GetWallets", "invalid limit", serr.ErrInvalidInput))
			return
		}

		if parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if o := c.Query("offset"); o != "" {
		parsedOffset, err := strconv.Atoi(o)
		if err != nil {
			handleError(c, serr.ValidationErr("handler.GetWallets", "invalid offset", serr.ErrInvalidInput))
			return
		}

		if parsedOffset > 0 {
			offset = parsedOffset
		}
	}
	wallets, err := a.walletService.GetWalletsWithPagination(c.Request.Context(), limit, offset)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, wallets)
}

// GetWallet godoc
// @Summary Get a specific wallet by user ID
// @Description Retrieves the wallet of a specific user by their user ID for admins.
// @Tags Admin Wallets
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} entity.Wallet "Wallet retrieved successfully"
// @Failure 400 {object} Error "Invalid user ID"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 404 {object} Error "Wallet not found"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/wallets/users/{id} [get]
func (a *AdminWalletHandler) GetWallet(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handleError(c, serr.ValidationErr("handler.GetWallet", "invalid user ID", serr.ErrInvalidInput))
		return
	}

	wallet, err := a.walletService.GetUserWallet(c.Request.Context(), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, wallet)
}
