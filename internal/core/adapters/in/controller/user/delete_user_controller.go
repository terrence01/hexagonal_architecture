package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"hexagonal-architecture/internal/core/ports/in"
)

// DeleteUserController handles HTTP requests for deleting users
type DeleteUserController struct {
	deleteUserUseCase in.DeleteUserUseCase
}

// NewDeleteUserController creates a new DeleteUserController
func NewDeleteUserController(deleteUserUseCase in.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{
		deleteUserUseCase: deleteUserUseCase,
	}
}

// HandleDeleteUser handles deleting a user by ID
func (c *DeleteUserController) HandleDeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	if err := c.deleteUserUseCase.Execute(ctx.Request.Context(), userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
