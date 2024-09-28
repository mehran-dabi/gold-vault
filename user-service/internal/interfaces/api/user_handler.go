package api

import (
	"net/http"

	"goldvault/user-service/internal/core/application/services"
	"goldvault/user-service/internal/interfaces/api/dto"
	"goldvault/user-service/pkg/serr"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetProfile godoc
// @Summary Get user profile
// @Description Retrieves the profile of a user by their ID.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} entity.User "User profile retrieved successfully"
// @Failure 400 {object} Error "Invalid user ID"
// @Failure 404 {object} Error "User not found"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /users/me [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	// get the user ID from the context
	userID := c.GetInt64("user_id")
	if userID == 0 {
		handleError(c, serr.ValidationErr("handler.UpdateProfile",
			"invalid user ID", serr.ErrInvalidInput))
		return
	}

	user, err := h.userService.GetProfile(c.Request.Context(), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Updates the profile of a user by their ID.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.UpdateUserRequest true "Updated user data"
// @Success 200 {object} nil "User profile updated successfully"
// @Failure 400 {object} Error "Invalid user ID or input data"
// @Failure 404 {object} Error "User not found"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /users/me [patch]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	// get the user ID from the context
	userID := c.GetInt64("user_id")
	if userID == 0 {
		handleError(c, serr.ValidationErr("handler.UpdateProfile",
			"invalid user ID", serr.ErrInvalidInput))
		return
	}

	var updatedUser dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		handleError(c, serr.ValidationErr("handler.UpdateProfile",
			"invalid input", serr.ErrInvalidInput))
		return
	}

	// use the user ID from the URL path to ensure the user is updating their own profile
	updatedUser.ID = userID

	err := h.userService.UpdateUser(c.Request.Context(), &updatedUser)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
