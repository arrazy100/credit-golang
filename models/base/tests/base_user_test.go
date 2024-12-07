package models_tests

import (
	"credit/models/base"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBaseUser_SetUser(t *testing.T) {
	baseUser := &base.BaseUser{}

	userID := uuid.New()

	baseUser.SetUser(userID)

	assert.Equal(t, baseUser.UserID, userID)
}
