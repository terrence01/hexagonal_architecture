package getusercontroller

import (
	"hexagonal-architecture/internal/core/adapters/in/http/dto"
	"hexagonal-architecture/internal/corev2/application/user/getuserservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserController handles HTTP requests for retrieving users
type GetUserController struct {
	getUserService getuserservice.GetUserService
}

// NewGetUserController creates a new GetUserController
func NewGetUserController(getUserService getuserservice.GetUserService) *GetUserController {
	return &GetUserController{
		getUserService: getUserService,
	}
}

// HandleGetUser handles retrieving a user by ID
func (c *GetUserController) HandleGetUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := c.getUserService.Execute(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	resp := dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	ctx.JSON(http.StatusOK, resp)
}
