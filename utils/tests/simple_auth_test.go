package utils_tests

import (
	"credit/models/enums"
	"credit/utils"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSimpleAuth_GenerateAndParseToken(t *testing.T) {
	userID := uuid.MustParse("e87afcf0-72f8-4531-94d1-23936f29a0e1")
	role := enums.Admin

	token, err := utils.GenerateToken(userID, role)
	assert.NoError(t, err)

	decoded, err := utils.ValidateTokenAndParseData(token)
	assert.NoError(t, err)

	assert.Equal(t, decoded.UserID, userID)
	assert.Equal(t, decoded.Role, role)
}

func TestSimpleAuth_ValidateRole(t *testing.T) {
	simpleAuth := utils.SimpleAuth{
		UserID: uuid.New(),
		Role:   enums.Debtor,
	}

	err := simpleAuth.ValidateRole(enums.Admin)
	assert.Error(t, err)

	err = simpleAuth.ValidateRole(enums.Debtor)
	assert.NoError(t, err)
}
