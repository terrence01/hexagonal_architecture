package user

import (
	"hexagonal-architecture/internal/core/adapters/in/http/dto"
	"net/http"

	"github.com/gin-gonic/gin"

	"hexagonal-architecture/internal/core/ports/in"
)

// GetUserController handles HTTP requests for retrieving users
type GetUserController struct {
	getUserUseCase in.GetUserUseCase
}

// NewGetUserController creates a new GetUserController
func NewGetUserController(getUserUseCase in.GetUserUseCase) *GetUserController {
	return &GetUserController{
		getUserUseCase: getUserUseCase,
	}
}

// HandleGetUser handles retrieving a user by ID
func (c *GetUserController) HandleGetUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := c.getUserUseCase.Execute(ctx.Request.Context(), userID)
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
