package domain

import (
	"time"
)

// User represents the user entity in our domain
type User struct {
	ID        string
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser creates a new User entity with the given properties
func NewUser(username, email string) *User {
	now := time.Now()
	return &User{
		ID:        generateID(), // In a real application, use a proper ID generation method
		Username:  username,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// generateID is a placeholder for ID generation
func generateID() string {
	// In a real application, use UUID or other ID generation method
	return time.Now().Format("20060102150405")
}
