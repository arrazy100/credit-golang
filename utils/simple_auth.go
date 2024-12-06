package utils

import (
	"credit/models/enums"
	"encoding/base64"
	"encoding/json"

	"github.com/google/uuid"
)

type SimpleAuth struct {
	UserID uuid.UUID      `json:"id"`
	Role   enums.UserRole `json:"role"`
}

func GenerateToken(userID uuid.UUID, role enums.UserRole) (string, error) {
	dummyAuth := &SimpleAuth{
		UserID: userID,
		Role:   role,
	}

	jsonData, err := json.Marshal(dummyAuth)
	if err != nil {
		return "", err
	}

	token := base64.StdEncoding.EncodeToString(jsonData)

	return token, nil
}

func ValidateTokenAndParseData(token string) (*SimpleAuth, error) {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	var simpleAuth SimpleAuth
	err = json.Unmarshal(decodedToken, &simpleAuth)
	if err != nil {
		return nil, err
	}

	return &simpleAuth, nil
}