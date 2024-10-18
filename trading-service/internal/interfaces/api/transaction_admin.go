package api

import (
	"net/http"
	"strconv"
	"time"

	"goldvault/trading-service/internal/core/application/services"
	"goldvault/trading-service/pkg/serr"

	"github.com/gin-gonic/gin"
)

type TransactionsAdminHandler struct {
	transactionService *services.TransactionService
}

func NewTransactionAdminHandler(transactionService *services.TransactionService) *TransactionsAdminHandler {
	return &TransactionsAdminHandler{
		transactionService: transactionService,
	}
}

// GetTransactions godoc
// @Summary Get all transactions with pagination
// @Description Retrieves a paginated list of all transactions for admins.
// @Tags Admin Transactions
// @Accept json
// @Produce json
// @Param limit query int false "Limit the number of results (default is 10)"
// @Param offset query int false "Offset for pagination (default is 0)"
// @Success 200 {object} map[string][]entity.Transaction "Transactions retrieved successfully"
// @Failure 400 {object} Error "Invalid limit or offset"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/transactions [get]
func (t *TransactionsAdminHandler) GetTransactions(c *gin.Context) {
	limit := 10
	offset := 0

	if l := c.Query("limit"); l != "" {
		parsedLimit, err := strconv.Atoi(l)
		if err != nil {
			handleError(c, serr.ValidationErr("TransactionsAdminHandler.GetTransactions", err.Error(), serr.ErrInvalidInput))
			return
		}
		limit = parsedLimit
	}

	if o := c.Query("offset"); o != "" {
		parsedOffset, err := strconv.Atoi(o)
		if err != nil {
			handleError(c, serr.ValidationErr("TransactionsAdminHandler.GetTransactions", err.Error(), serr.ErrInvalidInput))
			return
		}
		offset = parsedOffset
	}

	txs, err := t.transactionService.GetTransactions(c.Request.Context(), int64(limit), int64(offset))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": txs})
}

// GetUserTransactions godoc
// @Summary Get transactions for a specific user
// @Description Retrieves a list of transactions for a specific user by their ID.
// @Tags Admin Transactions
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {object} map[string]interface{} "User transactions retrieved successfully"
// @Failure 400 {object} Error "Invalid user ID"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 404 {object} Error "Transactions not found"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/transactions/{userID} [get]
func (t *TransactionsAdminHandler) GetUserTransactions(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		handleError(c, serr.ValidationErr("TransactionsAdminHandler.GetUserTransactions", err.Error(), serr.ErrInvalidInput))
		return
	}

	txs, err := t.transactionService.GetUserTransactions(c.Request.Context(), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": txs})
}

// GetSingleDaySummary godoc
// @Summary Get single day transaction summary
// @Description Retrieves a summary of transactions for a specific asset type on a given date.
// @Tags Admin Transactions
// @Produce json
// @Param date query string true "Date in YYYY-MM-DD format"
// @Param assetType query string true "Asset type"
// @Success 200 {object} map[string]interface{} "Summary of transactions for the day"
// @Failure 400 {object} Error "Invalid date format or missing parameters"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/transactions/summary/single-day [get]
func (t *TransactionsAdminHandler) GetSingleDaySummary(c *gin.Context) {
	dateStr := c.Query("date")
	assetType := c.Query("assetType")

	// Convert the string into a time.Time object
	date, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	summary, err := t.transactionService.GetSingleDaySummary(c.Request.Context(), date, assetType)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"summary": summary})
}

// GetTotalSummary godoc
// @Summary Get total transaction summary
// @Description Retrieves the total summary of all transactions for a specific asset type.
// @Tags Admin Transactions
// @Produce json
// @Param assetType query string true "Asset type"
// @Success 200 {object} map[string]interface{} "Total summary of transactions"
// @Failure 400 {object} Error "Invalid asset type or missing parameters"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/transactions/summary/total [get]
func (t *TransactionsAdminHandler) GetTotalSummary(c *gin.Context) {
	assetType := c.Query("assetType")

	summary, err := t.transactionService.GetTotalSummary(c.Request.Context(), assetType)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"summary": summary})
}
