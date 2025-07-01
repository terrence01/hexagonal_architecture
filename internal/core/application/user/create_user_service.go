package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"hexagonal-architecture/internal/core/domain"
	"hexagonal-architecture/internal/core/ports/in/command"
	"hexagonal-architecture/internal/core/ports/out"
)

// CreateUserService implements the CreateUserUseCase interface
type CreateUserService struct {
	userRepository domain.UserRepository
	emailSender    out.EmailSender
}

// NewCreateUserService creates a new CreateUserService
func NewCreateUserService(userRepository domain.UserRepository, emailSender out.EmailSender) *CreateUserService {
	return &CreateUserService{
		userRepository: userRepository,
		emailSender:    emailSender,
	}
}

// Execute creates a new user
func (s *CreateUserService) Execute(ctx context.Context, cmd command.CreateUserCommand) (*domain.User, error) {
	// 檢查電子郵件是否已被使用
	existingUser, err := s.userRepository.FindUserByEmail(ctx, cmd.Email)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		return nil, fmt.Errorf("檢查電子郵件時出錯: %w", err)
	}
	if existingUser != nil {
		return nil, ErrEmailAlreadyUsed
	}

	// 創建新用戶
	user := &domain.User{
		ID:        uuid.New().String(),
		Username:  cmd.Username,
		Email:     cmd.Email,
		CreatedAt: time.Now(),
	}

	// 保存用戶
	err = s.userRepository.SaveUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("保存用戶時出錯: %w", err)
	}

	// 發送歡迎電子郵件
	go func() {
		emailCtx := context.Background()
		subject := "歡迎加入我們的平台"
		body := fmt.Sprintf("親愛的 %s，\n\n感謝您註冊我們的服務！您的帳戶已成功創建。\n\n祝您使用愉快！", user.Username)

		err := s.emailSender.Send(emailCtx, user.Email, subject, body)
		if err != nil {
			// 在實際應用中，這裡應該記錄錯誤
			fmt.Printf("發送歡迎電子郵件失敗: %v\n", err)
		}
	}()

	return user, nil
}
