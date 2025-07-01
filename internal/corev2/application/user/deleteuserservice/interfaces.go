package deleteuserservice

import (
	"context"
	"hexagonal-architecture/internal/corev2/domain"
)

type UserRepository interface {
	// DeleteUser removes a user from the database by ID
	DeleteUser(ctx context.Context, userID string) error

	// FindUserByID retrieves a user from the database by ID
	FindUserByID(ctx context.Context, userID string) (*domain.User, error)
}

type EmailSender interface {
	// Send sends an email with the given details
	Send(ctx context.Context, to, subject, body string) error
}

type DeleteUserService interface {
	// Execute deletes a user by their ID
	Execute(ctx context.Context, userID string) error
}
