package response

import "credit/utils"

type LoginResponse struct {
	Token string           `json:"token"`
	User  utils.SimpleAuth `json:"user"`
}
