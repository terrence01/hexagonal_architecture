package user

import (
	"hexagonal-architecture/internal/core/domain"
	"hexagonal-architecture/internal/core/ports/out"
)

// UserService combines all user-related services
type UserService struct {
	CreateUserService
	DeleteUserService
	GetUserService
}

// NewUserService creates a new UserService that embeds all specialized services
func NewUserService(userRepository domain.UserRepository, emailSender out.EmailSender) *UserService {
	return &UserService{
		CreateUserService: *NewCreateUserService(userRepository, emailSender),
		DeleteUserService: *NewDeleteUserService(userRepository, emailSender),
		GetUserService:    *NewGetUserService(userRepository, emailSender),
	}
}
