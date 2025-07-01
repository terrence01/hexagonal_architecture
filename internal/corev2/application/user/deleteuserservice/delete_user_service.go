package deleteuserservice

import (
	"context"
	"fmt"

	"hexagonal-architecture/internal/core/ports/out"
)

// Service implements the DeleteUserUseCase interface
type Service struct {
	userRepository UserRepository
	emailSender    out.EmailSender
}

// NewService creates a new Service
func NewService(userRepository UserRepository, emailSender out.EmailSender) *Service {
	return &Service{
		userRepository: userRepository,
		emailSender:    emailSender,
	}
}

// Execute deletes a user by ID
func (s *Service) Execute(ctx context.Context, userID string) error {
	// First get user information to send email
	user, err := s.userRepository.FindUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("error finding user: %w", err)
	}

	// Delete user
	err = s.userRepository.DeleteUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	// Send account deletion notification email
	go func() {
		emailCtx := context.Background()
		subject := "Your account has been deleted"
		body := fmt.Sprintf("Dear %s,\n\nYour account has been successfully deleted. Thank you for using our service.\n\nIf this was not your action, please contact our customer support immediately.", user.Username)

		err := s.emailSender.Send(emailCtx, user.Email, subject, body)
		if err != nil {
			// In a real application, this error should be logged
			fmt.Printf("Failed to send account deletion notification email: %v\n", err)
		}
	}()

	return nil
}
