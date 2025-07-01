package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteUserController handles HTTP requests for deleting users
// HandleDeleteUser handles deleting a user by ID
func (c *Controller) HandleDeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	if err := c.userService.DeleteUser(ctx.Request.Context(), userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
