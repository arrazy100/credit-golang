package validations_test

import (
	validations "credit/validations"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestPayload struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Age   int    `validate:"required,gte=18"`
}

func TestValidateStruct(t *testing.T) {
	validPayload := TestPayload{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Age:   25,
	}

	validErrors := validations.ValidateStruct(validPayload)
	assert.Empty(t, validErrors, "Expected no validation errors for valid payload")

	invalidPayload := TestPayload{
		Name:  "",
		Email: "invalid-email",
		Age:   15,
	}

	invalidErrors := validations.ValidateStruct(invalidPayload)
	invalidFields := invalidErrors.Fields
	assert.Len(t, invalidFields, 3, "Expected 3 validation errors for invalid payload")

	assert.Equal(t, "Name", *invalidFields[0].Field)
	assert.Equal(t, "required", *invalidFields[0].Tag)
	assert.Empty(t, invalidFields[0].Value)

	assert.Equal(t, "Email", *invalidFields[1].Field)
	assert.Equal(t, "email", *invalidFields[1].Tag)
	assert.Empty(t, invalidFields[1].Value)

	assert.Equal(t, "Age", *invalidFields[2].Field)
	assert.Equal(t, "gte", *invalidFields[2].Tag)
	assert.Equal(t, "18", invalidFields[2].Value)
}

func TestConvert(t *testing.T) {
	testError := errors.New("an unexpected error occurred")
	convertedErrors := validations.Convert(testError)

	assert.Equal(t, "an unexpected error occurred", convertedErrors.Message)

	nilErrors := validations.Convert(nil)
	assert.Empty(t, nilErrors, "Expected no errors for nil input")
}
