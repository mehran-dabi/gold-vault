package api

import (
	"net/http"

	"goldvault/wallet-service/internal/core/application/services"
	"goldvault/wallet-service/pkg/serr"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	walletService      *services.WalletService
	assetService       *services.AssetService
	transactionService *services.TransactionService
}

func NewWalletHandler(
	walletService *services.WalletService,
	assetService *services.AssetService,
	transactionService *services.TransactionService,
) *WalletHandler {
	return &WalletHandler{
		walletService:      walletService,
		assetService:       assetService,
		transactionService: transactionService,
	}
}

// GetUserWallet godoc
// @Summary Get user wallet
// @Description Retrieves the wallet information of the authenticated user by their ID.
// @Tags Wallets
// @Accept json
// @Produce json
// @Success 200 {object} entity.Wallet "Wallet retrieved successfully"
// @Failure 400 {object} Error "Invalid user ID"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 404 {object} Error "Wallet not found"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /wallets [get]
func (w *WalletHandler) GetUserWallet(c *gin.Context) {
	// Extract the userID from the JWT (set in the context by JWT middleware)
	userID := c.GetInt64("user_id")

	wallet, err := w.walletService.GetUserWallet(c.Request.Context(), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, wallet)
}

// GetWalletTransactions godoc
// @Summary Get wallet transactions
// @Description Retrieves all transactions for a specific wallet by wallet ID.
// @Tags Wallets
// @Accept json
// @Produce json
// @Success 200 {array} entity.Transaction "List of transactions"
// @Failure 400 {object} Error "Invalid wallet ID"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 404 {object} Error "Wallet not found"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /wallets/transactions [get]
func (w *WalletHandler) GetWalletTransactions(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		handleError(c, serr.ValidationErr("handler.GetWalletTransactions",
			"invalid user ID", serr.ErrUnAuthorized))
		return
	}

	transactions, err := w.transactionService.GetWalletTransactions(c.Request.Context(), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, transactions)
}
