package user

import (
	"github.com/gin-gonic/gin"

	"hexagonal-architecture/internal/core/ports/in"
)

// UserController combines all user-related controllers
type UserController struct {
	CreateUserController *CreateUserController
	GetUserController    *GetUserController
	DeleteUserController *DeleteUserController
}

// NewUserController creates a new UserController that combines all user-related controllers
func NewUserController(
	createUserUseCase in.CreateUserUseCase,
	deleteUserUseCase in.DeleteUserUseCase,
	getUserUseCase in.GetUserUseCase,
) *UserController {
	return &UserController{
		CreateUserController: NewCreateUserController(createUserUseCase),
		GetUserController:    NewGetUserController(getUserUseCase),
		DeleteUserController: NewDeleteUserController(deleteUserUseCase),
	}
}

// HandleCreateUser delegates to the CreateUserController
func (c *UserController) HandleCreateUser(ctx *gin.Context) {
	c.CreateUserController.HandleCreateUser(ctx)
}

// HandleGetUser delegates to the GetUserController
func (c *UserController) HandleGetUser(ctx *gin.Context) {
	c.GetUserController.HandleGetUser(ctx)
}

// HandleDeleteUser delegates to the DeleteUserController
func (c *UserController) HandleDeleteUser(ctx *gin.Context) {
	c.DeleteUserController.HandleDeleteUser(ctx)
}
