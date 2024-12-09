package response

import "github.com/google/uuid"

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
}
