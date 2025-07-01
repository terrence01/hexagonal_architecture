package dto

// CreateUserRequest represents the request data for creating a user
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}
