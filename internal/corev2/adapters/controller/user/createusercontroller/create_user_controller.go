package createusercontroller

import (
	"hexagonal-architecture/internal/corev2/adapters/http/dto"
	"hexagonal-architecture/internal/corev2/application/user/createuserservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller handles HTTP requests for creating users
type Controller struct {
	createUserService createuserservice.CreateUserService
}

// NewCreateUserController creates a new Controller
func NewCreateUserController(createUserService createuserservice.CreateUserService) *Controller {
	return &Controller{
		createUserService: createUserService,
	}
}

// HandleCreateUser handles the creation of a new user
func (c *Controller) HandleCreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	cmd := createuserservice.Command{
		Username: req.Username,
		Email:    req.Email,
	}

	user, err := c.createUserService.Execute(ctx.Request.Context(), cmd)
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
