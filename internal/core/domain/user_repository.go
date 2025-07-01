package domain

import (
	"context"
)

// UserRepository defines the repository interface for User entity
type UserRepository interface {
	// SaveUser persists a user to the database
	SaveUser(ctx context.Context, user *User) error
	
	// DeleteUser removes a user from the database by ID
	DeleteUser(ctx context.Context, userID string) error
	
	// FindUserByID retrieves a user from the database by ID
	FindUserByID(ctx context.Context, userID string) (*User, error)
	
	// FindUserByEmail retrieves a user from the database by email
	FindUserByEmail(ctx context.Context, email string) (*User, error)
}
