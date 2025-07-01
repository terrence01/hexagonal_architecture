package email

import (
	"context"
)

// EmailSender implements the EmailSender interface
type EmailSender struct{}

// NewEmailSender creates a new EmailSender
func NewEmailSender() *EmailSender {
	return &EmailSender{}
}

// Send implements the EmailSender interface
func (s *EmailSender) Send(ctx context.Context, to, subject, body string) error {
	// 簡單實現，直接返回 nil
	return nil
}
