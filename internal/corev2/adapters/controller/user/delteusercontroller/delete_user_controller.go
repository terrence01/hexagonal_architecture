package delteusercontroller

import (
	"hexagonal-architecture/internal/corev2/application/user/deleteuserservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller handles HTTP requests for deleting users
type Controller struct {
	deleteUserService deleteuserservice.DeleteUserService
}

// NewController creates a new Controller
func NewController(deleteUserService deleteuserservice.DeleteUserService) *Controller {
	return &Controller{
		deleteUserService: deleteUserService,
	}
}

// HandleDeleteUser handles deleting a user by ID
func (c *Controller) HandleDeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	if err := c.deleteUserService.Execute(ctx.Request.Context(), userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
