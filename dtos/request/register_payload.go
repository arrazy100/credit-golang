package request

type RegisterPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=25"`
}
