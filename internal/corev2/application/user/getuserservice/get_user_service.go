package getuserservice

import (
	"context"
	"fmt"

	"hexagonal-architecture/internal/corev2/domain"
)

// Service implements the GetUserUseCase interface
type Service struct {
	userRepository UserRepository
}

// NewService creates a new Service
func NewService(userRepository UserRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

// Execute retrieves a user by ID
func (s *Service) Execute(ctx context.Context, userID string) (*domain.User, error) {
	user, err := s.userRepository.FindUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error finding user: %w", err)
	}
	return user, nil
}
