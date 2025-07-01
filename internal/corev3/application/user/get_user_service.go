package user

import (
	"context"
	"fmt"

	"hexagonal-architecture/internal/corev3/domain"
)

func (s *Service) GetUser(ctx context.Context, userID string) (*domain.User, error) {
	user, err := s.userRepository.FindUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("查找用戶時出錯: %w", err)
	}
	return user, nil
}
