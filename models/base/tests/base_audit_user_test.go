package models_tests

import (
	"credit/models/base"
	"credit/models/enums"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBaseAuditUser_SetCreated(t *testing.T) {
	baseAudit := &base.BaseAuditUser{}
	user := base.User{
		ID:       uuid.New(),
		Email:    "user@user.com",
		Password: "user123",
		Role:     enums.Debtor,
	}

	baseAudit.SetAuditCreated(time.Now(), user)

	assert.Equal(t, baseAudit.CreatedByID, user.ID)
}

func TestBaseAuditUser_SetUpdated(t *testing.T) {
	baseAudit := &base.BaseAuditUser{}
	user := base.User{
		ID:       uuid.New(),
		Email:    "user@user.com",
		Password: "user123",
		Role:     enums.Debtor,
	}

	baseAudit.SetAuditUpdated(time.Now(), user)

	assert.Equal(t, baseAudit.UpdatedByID, user.ID)
}
