package models_tests

import (
	"credit/models/base"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBaseAuditUser_SetCreated(t *testing.T) {
	baseAudit := &base.BaseAuditUser{}

	userID := uuid.New()

	baseAudit.SetAuditCreated(time.Now(), userID)

	assert.Equal(t, baseAudit.CreatedByID, userID)
}

func TestBaseAuditUser_SetUpdated(t *testing.T) {
	baseAudit := &base.BaseAuditUser{}

	userID := uuid.New()

	baseAudit.SetAuditUpdated(time.Now(), userID)

	assert.Equal(t, baseAudit.UpdatedByID, &userID)
}
