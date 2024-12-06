package utils_tests

import (
	"credit/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtils_HashAndCheckPassword(t *testing.T) {
	password := "password123"

	hashed, err := utils.HashPassword(password)
	assert.NoError(t, err)

	valid := utils.CheckPasswordHash(password, hashed)
	assert.Equal(t, valid, true)
}
