package interfaces

import (
	"context"
	"hexagonal-architecture/internal/corev3/application/user/command"
	"hexagonal-architecture/internal/corev3/domain"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, userID string) error
	FindUserByID(ctx context.Context, userID string) (*domain.User, error)
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type EmailSender interface {
	// Send sends an email with the given details
	Send(ctx context.Context, to, subject, body string) error
}

type UserService interface {
	CreateUser(ctx context.Context, cmd command.CreateUserCommand) (*domain.User, error)
	DeleteUser(ctx context.Context, userID string) error
	GetUser(ctx context.Context, userID string) (*domain.User, error)
}
