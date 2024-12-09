package response

import (
	"github.com/google/uuid"
)

type RegisterResponse struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Role   string    `json:"role"`
}
