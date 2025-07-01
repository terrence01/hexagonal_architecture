package in

import (
	"context"

	"hexagonal-architecture/internal/core/domain"
	"hexagonal-architecture/internal/core/ports/in/command"
)

// CreateUserUseCase defines the port for creating a user
type CreateUserUseCase interface {
	// Execute creates a new user with the given details
	Execute(ctx context.Context, cmd command.CreateUserCommand) (*domain.User, error)
}

// DeleteUserUseCase defines the port for deleting a user
type DeleteUserUseCase interface {
	// Execute deletes a user by their ID
	Execute(ctx context.Context, userID string) error
}

// GetUserUseCase defines the port for retrieving a user
type GetUserUseCase interface {
	// Execute retrieves a user by their ID
	Execute(ctx context.Context, userID string) (*domain.User, error)
}
