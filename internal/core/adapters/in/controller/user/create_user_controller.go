package user

import (
	"hexagonal-architecture/internal/core/adapters/in/http/dto"
	"net/http"

	"github.com/gin-gonic/gin"

	"hexagonal-architecture/internal/core/ports/in"
	"hexagonal-architecture/internal/core/ports/in/command"
)

// CreateUserController handles HTTP requests for creating users
type CreateUserController struct {
	createUserUseCase in.CreateUserUseCase
}

// NewCreateUserController creates a new CreateUserController
func NewCreateUserController(createUserUseCase in.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{
		createUserUseCase: createUserUseCase,
	}
}

// HandleCreateUser handles the creation of a new user
func (c *CreateUserController) HandleCreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	cmd := command.CreateUserCommand{
		Username: req.Username,
		Email:    req.Email,
	}

	user, err := c.createUserUseCase.Execute(ctx.Request.Context(), cmd)
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
