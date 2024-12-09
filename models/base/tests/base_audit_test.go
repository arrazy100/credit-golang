package models_tests

import (
	"credit/models/base"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBaseAudit_SetCreated(t *testing.T) {
	baseAudit := &base.BaseAudit{}
	currentTime := time.Now()

	baseAudit.SetAuditCreated(currentTime)

	assert.Equal(t, baseAudit.CreatedAt, currentTime.UTC())
}

func TestBaseAudit_SetUpdated(t *testing.T) {
	baseAudit := &base.BaseAudit{}
	currentTime := time.Now()

	baseAudit.SetAuditUpdated(currentTime)

	assert.Equal(t, baseAudit.UpdatedAt, currentTime.UTC())
}
