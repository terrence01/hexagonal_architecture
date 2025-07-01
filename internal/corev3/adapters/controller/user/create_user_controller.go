package user

import (
	"hexagonal-architecture/internal/core/adapters/in/http/dto"
	"hexagonal-architecture/internal/corev3/application/user/command"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUserController handles HTTP requests for creating users

// HandleCreateUser handles the creation of a new user
func (c *Controller) HandleCreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	cmd := command.CreateUserCommand{
		Username: req.Username,
		Email:    req.Email,
	}

	user, err := c.userService.CreateUser(ctx.Request.Context(), cmd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	ctx.JSON(http.StatusCreated, resp)
}
