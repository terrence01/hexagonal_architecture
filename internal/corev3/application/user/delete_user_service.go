package user

import (
	"context"
	"fmt"
)

func (s *Service) DeleteUser(ctx context.Context, userID string) error {
	// 先獲取用戶信息，以便發送電子郵件
	user, err := s.userRepository.FindUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("查找用戶時出錯: %w", err)
	}

	// 刪除用戶
	err = s.userRepository.DeleteUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("刪除用戶時出錯: %w", err)
	}

	// 發送帳戶刪除通知電子郵件
	go func() {
		emailCtx := context.Background()
		subject := "您的帳戶已被刪除"
		body := fmt.Sprintf("親愛的 %s，\n\n您的帳戶已被成功刪除。感謝您使用我們的服務。\n\n如果這不是您的操作，請立即聯繫我們的客戶支持。", user.Username)

		err := s.emailSender.Send(emailCtx, user.Email, subject, body)
		if err != nil {
			// 在實際應用中，這裡應該記錄錯誤
			fmt.Printf("發送帳戶刪除通知電子郵件失敗: %v\n", err)
		}
	}()

	return nil
}
