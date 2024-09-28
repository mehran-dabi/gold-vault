package api

import (
	"net/http"

	"goldvault/trading-service/internal/core/application/services"
	"goldvault/trading-service/pkg/serr"

	"github.com/gin-gonic/gin"
)

type TransactionsHandler struct {
	transactionService *services.TransactionService
}

func NewTransactionsHandler(transactionService *services.TransactionService) *TransactionsHandler {
	return &TransactionsHandler{
		transactionService: transactionService,
	}
}

// GetUserTransactions godoc
// @Summary Get transactions for the authenticated user
// @Description Retrieves a list of all transactions for the authenticated user.
// @Tags User Transactions
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]entity.Transaction "User transactions retrieved successfully"
// @Failure 400 {object} Error "Invalid user ID"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /trades/transactions [get]
func (t *TransactionsHandler) GetUserTransactions(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		handleError(c, serr.ValidationErr("TransactionsHandler.GetUserTransactions", "invalid user ID", serr.ErrInvalidInput))
		return
	}

	txs, err := t.transactionService.GetUserTransactions(c.Request.Context(), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": txs})
}
