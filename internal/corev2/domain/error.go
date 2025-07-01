package domain

import "errors"

var (
	// ErrUserNotFound is returned when a user is not found
	ErrUserNotFound = errors.New("user not found")

	// ErrEmailAlreadyUsed is returned when an email is already in use
	ErrEmailAlreadyUsed = errors.New("email already in use")
)
