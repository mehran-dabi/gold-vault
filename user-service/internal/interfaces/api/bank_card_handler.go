package api

import (
	"goldvault/user-service/internal/core/application/services"
	"goldvault/user-service/internal/interfaces/api/dto"
	"goldvault/user-service/pkg/serr"

	"github.com/gin-gonic/gin"
)

type BankCardHandler struct {
	bankCardService *services.BankCardService
}

func NewBankCardHandler(bankCardService *services.BankCardService) *BankCardHandler {
	return &BankCardHandler{
		bankCardService: bankCardService,
	}
}

// AddUserBankCard godoc
// @Summary Add a bank card for the user
// @Description Adds a new bank card for the authenticated user.
// @Tags User Bank Cards
// @Accept json
// @Produce json
// @Param request body dto.AddUserBankCardRequest true "Add bank card request"
// @Success 200 {object} map[string]string "Bank card added successfully"
// @Failure 400 {object} Error "Invalid input"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /users/bank-cards [post]
func (b *BankCardHandler) AddUserBankCard(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		handleError(c, serr.ValidationErr("handler.UpdateProfile",
			"invalid user ID", serr.ErrInvalidInput))
		return
	}

	var req dto.AddUserBankCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, serr.ValidationErr("AddUserBankCard", err.Error(), serr.ErrInvalidInput))
		return
	}

	if err := req.Validate(); err != nil {
		handleError(c, serr.ValidationErr("AddUserBankCard", err.Error(), serr.ErrInvalidInput))
		return
	}

	err := b.bankCardService.AddUserBankCard(c, userID, req.CardNumber)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "Bank card added successfully"})
}

// GetUserBankCards godoc
// @Summary Get bank cards for the user
// @Description Retrieves a list of all bank cards for the authenticated user.
// @Tags User Bank Cards
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "List of user's bank cards"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /users/bank-cards [get]
func (b *BankCardHandler) GetUserBankCards(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		handleError(c, serr.ValidationErr("handler.UpdateProfile",
			"invalid user ID", serr.ErrInvalidInput))
		return
	}

	bankCards, err := b.bankCardService.GetUserBankCards(c, userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{"bank_cards": bankCards})
}
