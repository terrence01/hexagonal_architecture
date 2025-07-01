package user

import (
	"hexagonal-architecture/internal/corev3/interfaces"
)

// Controller combines all user-related controllers

type Controller struct {
	userService interfaces.UserService
}

func NewUserController(userUseCase interfaces.UserService) *Controller {
	return &Controller{
		userService: userUseCase,
	}
}
