package user

import (
	"context"
	"fmt"

	"hexagonal-architecture/internal/core/domain"
	"hexagonal-architecture/internal/core/ports/out"
)

// GetUserService implements the GetUserUseCase interface
type GetUserService struct {
	userRepository domain.UserRepository
	emailSender    out.EmailSender
}

// NewGetUserService creates a new GetUserService
func NewGetUserService(userRepository domain.UserRepository, emailSender out.EmailSender) *GetUserService {
	return &GetUserService{
		userRepository: userRepository,
		emailSender:    emailSender,
	}
}

// Execute retrieves a user by ID
func (s *GetUserService) Execute(ctx context.Context, userID string) (*domain.User, error) {
	user, err := s.userRepository.FindUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("查找用戶時出錯: %w", err)
	}
	return user, nil
}
