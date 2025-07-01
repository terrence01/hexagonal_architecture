package user

import (
	"hexagonal-architecture/internal/corev3/interfaces"
)

type Service struct {
	userRepository interfaces.UserRepository
	emailSender    interfaces.EmailSender
}

func NewService(userRepository interfaces.UserRepository, emailSender interfaces.EmailSender) *Service {
	return &Service{
		userRepository: userRepository,
		emailSender:    emailSender,
	}
}
