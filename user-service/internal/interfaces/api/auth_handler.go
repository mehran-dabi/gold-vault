package api

import (
	"net/http"

	"goldvault/user-service/internal/core/application/services"
	"goldvault/user-service/internal/interfaces/api/dto"
	"goldvault/user-service/pkg/serr"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

// GenerateOTP godoc
// @Summary      Sends a one-time password to the user's phone
// @Description  Sends a one-time password to the user's phone
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.GenerateOTPRequest true "Generate OTP Request"
// @Success      201  {object}  nil
// @Failure      400  {object}  Error
// @Failure      500  {object}  Error
// @Router       /auth/otp [post]
func (h *AuthHandler) GenerateOTP(c *gin.Context) {
	var req dto.GenerateOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, serr.ValidationErr("handler.GenerateOTP",
			"invalid input", serr.ErrInvalidInput))
		return
	}

	err := h.AuthService.RequestOTP(c.Request.Context(), &req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

// ValidateOTP godoc
// @Summary      Validates the one-time password sent to the user's phone
// @Description  Validates the one-time password sent to the user's phone
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.ValidateOTPRequest true "Validate OTP Request"
// @Success      201  {object}  nil
// @Failure      400  {object}  Error
// @Failure      500  {object}  Error
// @Router       /auth/otp/validate [post]
func (h *AuthHandler) ValidateOTP(c *gin.Context) {
	var req dto.ValidateOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, serr.ValidationErr("handler.ValidateOTP",
			"invalid input", serr.ErrInvalidInput))
		return
	}

	token, err := h.AuthService.VerifyOTPAndIssueToken(c.Request.Context(), &req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
