package user

import (
	"context"
	"errors"
	"fmt"
	"hexagonal-architecture/internal/corev3/application/user/command"
	"time"

	"github.com/google/uuid"

	"hexagonal-architecture/internal/corev3/domain"
)

// CreateUserService implements the CreateUserUseCase interface
// Execute creates a new user
func (s *Service) CreateUser(ctx context.Context, cmd command.CreateUserCommand) (*domain.User, error) {
	// Check if email is already in use
	existingUser, err := s.userRepository.FindUserByEmail(ctx, cmd.Email)
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		return nil, fmt.Errorf("error checking email: %w", err)
	}
	if existingUser != nil {
		return nil, domain.ErrEmailAlreadyUsed
	}

	// Create new user
	user := &domain.User{
		ID:        uuid.New().String(),
		Username:  cmd.Username,
		Email:     cmd.Email,
		CreatedAt: time.Now(),
	}

	// Save user
	err = s.userRepository.SaveUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error saving user: %w", err)
	}

	// Send welcome email
	go func() {
		emailCtx := context.Background()
		subject := "Welcome to our platform"
		body := fmt.Sprintf("Dear %s,\n\nThank you for registering with our service! Your account has been successfully created.\n\nEnjoy using our platform!", user.Username)

		err := s.emailSender.Send(emailCtx, user.Email, subject, body)
		if err != nil {
			// In a real application, this error should be logged
			fmt.Printf("Failed to send welcome email: %v\n", err)
		}
	}()

	return user, nil
}
