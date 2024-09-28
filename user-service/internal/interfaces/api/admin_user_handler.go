package api

import (
	"net/http"
	"strconv"

	"goldvault/user-service/internal/core/application/services"
	"goldvault/user-service/internal/interfaces/api/dto"
	"goldvault/user-service/pkg/serr"

	"github.com/gin-gonic/gin"
)

type AdminUserHandler struct {
	userService *services.UserService
}

func NewAdminUserHandler(userService *services.UserService) *AdminUserHandler {
	return &AdminUserHandler{userService: userService}
}

// GetUsers godoc
// @Summary Get list of users with pagination
// @Description Retrieves a paginated list of users for admins.
// @Tags Admin Users
// @Accept json
// @Produce json
// @Param limit query int false "Limit the number of results (default is 10)"
// @Param offset query int false "Offset for pagination (default is 0)"
// @Success 200 {array} entity.User "List of users retrieved successfully"
// @Failure 400 {object} Error "Invalid limit or offset"
// @Failure 401 {object} Error "Unauthorized"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/users [get]
func (a *AdminUserHandler) GetUsers(c *gin.Context) {
	limit := 10
	offset := 0

	if l := c.Query("limit"); l != "" {
		parsedLimit, err := strconv.Atoi(l)
		if err != nil {
			handleError(c, serr.ValidationErr("adminHandler.GetUsers", "invalid limit", serr.ErrInvalidInput))
			return
		}

		if parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if o := c.Query("offset"); o != "" {
		parsedOffset, err := strconv.Atoi(o)
		if err != nil {
			handleError(c, serr.ValidationErr("adminHandler.GetUsers", "invalid offset", serr.ErrInvalidInput))
			return
		}

		if parsedOffset > 0 {
			offset = parsedOffset
		}
	}

	users, err := a.userService.GetUsers(c.Request.Context(), int64(limit), int64(offset))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Retrieves the profile of a user by their ID.
// @Tags Admin Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} entity.User "User profile retrieved successfully"
// @Failure 400 {object} Error "Invalid user ID"
// @Failure 404 {object} Error "User not found"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/users/{id} [get]
func (a *AdminUserHandler) GetProfile(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		handleError(c, serr.ValidationErr("adminHandler.GetProfile",
			"invalid user ID", serr.ErrInvalidInput))
		return
	}

	user, err := a.userService.GetProfile(c.Request.Context(), userID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Updates the profile of a user by their ID.
// @Tags Admin Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body dto.UpdateUserRequest true "Updated user data"
// @Success 200 {object} nil "User profile updated successfully"
// @Failure 400 {object} Error "Invalid user ID or input data"
// @Failure 404 {object} Error "User not found"
// @Failure 500 {object} Error "Internal server error"
// @Security BearerAuth
// @Router /admin/users/{id} [patch]
func (a *AdminUserHandler) UpdateProfile(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		handleError(c, serr.ValidationErr("handler.UpdateProfile",
			"invalid user ID", serr.ErrInvalidInput))
		return
	}

	var req dto.AdminUpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, serr.ValidationErr("handler.UpdateProfile",
			"invalid input", serr.ErrInvalidInput))
		return
	}

	// use the user ID from the URL path to ensure the user is updating their own profile
	req.ID = userID

	err = a.userService.AdminUpdateUser(c.Request.Context(), &req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
