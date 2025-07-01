package createuserservice

import (
	"context"
	"hexagonal-architecture/internal/corev2/domain"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *domain.User) error
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type EmailSender interface {
	Send(ctx context.Context, to, subject, body string) error
}

type CreateUserService interface {
	Execute(ctx context.Context, cmd Command) (*domain.User, error)
}
