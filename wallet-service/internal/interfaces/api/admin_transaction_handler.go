package api

import (
	"net/http"
	"strconv"

	"goldvault/wallet-service/internal/core/application/services"
	"goldvault/wallet-service/pkg/serr"

	"github.com/gin-gonic/gin"
)

type AdminTransactionHandler struct {
	TransactionService *services.TransactionService
}

func NewAdminTransactionHandler(transactionService *services.TransactionService) *AdminTransactionHandler {
	return &AdminTransactionHandler{
		TransactionService: transactionService,
	}
}

// GetTransactions godoc
// @Summary Get all transactions with pagination
// @Description Retrieves a list of all transactions with pagination for admins.
// @Tags Admin Transactions
// @Accept json
// @Produce json
// @Param limit query int false "Limit the number of results"
// @Param offset query int false "Offset for pagination"
// @Success 200 {array} entity.Transaction "List of transactions retrieved successfully"
// @Failure 400 {object} Error "Invalid pagination parameters"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/transactions [get]
func (a *AdminTransactionHandler) GetTransactions(c *gin.Context) {
	limit := 10
	offset := 0

	if l := c.Query("limit"); l != "" {
		parsedLimit, err := strconv.Atoi(l)
		if err != nil {
			handleError(c, serr.ValidationErr("AdminTransactionHandler.GetTransactions", "invalid limit", serr.ErrInvalidInput))
			return
		}

		if parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if o := c.Query("offset"); o != "" {
		parsedOffset, err := strconv.Atoi(o)
		if err != nil {
			handleError(c, serr.ValidationErr("AdminTransactionHandler.GetTransactions", "invalid offset", serr.ErrInvalidInput))
			return
		}

		if parsedOffset > 0 {
			offset = parsedOffset
		}
	}

	transactions, err := a.TransactionService.GetTransactionsWithPagination(c.Request.Context(), limit, offset)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, transactions)
}
