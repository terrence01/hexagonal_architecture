package user

import (
	"hexagonal-architecture/internal/corev2/adapters/controller/user/createusercontroller"
	"hexagonal-architecture/internal/corev2/adapters/controller/user/delteusercontroller"
	"hexagonal-architecture/internal/corev2/adapters/controller/user/getusercontroller"
	"hexagonal-architecture/internal/corev2/application/user/createuserservice"
	"hexagonal-architecture/internal/corev2/application/user/deleteuserservice"
	"hexagonal-architecture/internal/corev2/application/user/getuserservice"

	"github.com/gin-gonic/gin"
)

// Controller combines all user-related controllers
type Controller struct {
	CreateUserController *createusercontroller.Controller
	GetUserController    *getusercontroller.GetUserController
	DeleteUserController *delteusercontroller.Controller
}

// NewController creates a new Controller that combines all user-related controllers
func NewController(
	createUserUseCase createuserservice.CreateUserService,
	deleteUserUseCase deleteuserservice.DeleteUserService,
	getUserUseCase getuserservice.GetUserService,
) *Controller {
	return &Controller{
		CreateUserController: createusercontroller.NewCreateUserController(createUserUseCase),
		GetUserController:    getusercontroller.NewGetUserController(getUserUseCase),
		DeleteUserController: delteusercontroller.NewController(deleteUserUseCase),
	}
}

// HandleCreateUser delegates to the CreateUserController
func (c *Controller) HandleCreateUser(ctx *gin.Context) {
	c.CreateUserController.HandleCreateUser(ctx)
}

// HandleGetUser delegates to the GetUserController
func (c *Controller) HandleGetUser(ctx *gin.Context) {
	c.GetUserController.HandleGetUser(ctx)
}

// HandleDeleteUser delegates to the DeleteUserController
func (c *Controller) HandleDeleteUser(ctx *gin.Context) {
	c.DeleteUserController.HandleDeleteUser(ctx)
}
