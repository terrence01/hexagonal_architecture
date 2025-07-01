package user

import (
	"hexagonal-architecture/internal/core/adapters/in/http/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) HandleGetUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := c.userService.GetUser(ctx.Request.Context(), userID)
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
