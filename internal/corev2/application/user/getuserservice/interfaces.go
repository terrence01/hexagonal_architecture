package getuserservice

import (
	"context"
	"hexagonal-architecture/internal/corev2/domain"
)

type UserRepository interface {
	// FindUserByID retrieves a user from the database by ID
	FindUserByID(ctx context.Context, userID string) (*domain.User, error)
}

type GetUserService interface {
	Execute(ctx context.Context, userID string) (*domain.User, error)
}
