package memory

import (
	"context"
	"sync"

	"hexagonal-architecture/internal/corev2/domain"
)

// UserRepository implements the UserRepository interface with an in-memory database
type UserRepository struct {
	users map[string]*domain.User
	mutex sync.RWMutex
}

// NewUserRepository creates a new in-memory UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*domain.User),
	}
}

// SaveUser persists a user to the in-memory database
func (r *UserRepository) SaveUser(ctx context.Context, user *domain.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.users[user.ID] = user
	return nil
}

// DeleteUser removes a user from the in-memory database by ID
func (r *UserRepository) DeleteUser(ctx context.Context, userID string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[userID]; !exists {
		return domain.ErrUserNotFound
	}

	delete(r.users, userID)
	return nil
}

// FindUserByID retrieves a user from the in-memory database by ID
func (r *UserRepository) FindUserByID(ctx context.Context, userID string) (*domain.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[userID]
	if !exists {
		return nil, domain.ErrUserNotFound
	}

	return user, nil
}

// FindUserByEmail retrieves a user from the in-memory database by email
func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, domain.ErrUserNotFound
}
