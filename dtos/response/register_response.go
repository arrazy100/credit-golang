package response

import (
	"credit/models/enums"

	"github.com/google/uuid"
)

type RegisterResponse struct {
	UserID uuid.UUID      `json:"user_id"`
	Email  string         `json:"email"`
	Role   enums.UserRole `json:"role"`
}
