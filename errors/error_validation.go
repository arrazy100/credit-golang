package custom_errors

import (
	"credit/utils"

	"github.com/go-playground/validator/v10"
)

type ErrorValidation struct {
	Fields  []ErrorField `json:"fields"`
	Message string       `json:"message"`
}

type ErrorField struct {
	Field *string `json:"field"`
	Tag   *string `json:"tag"`
	Value string  `json:"value"`
}

func toPointer(s string) *string {
	return &s
}

func ValidateStruct[T any](payload T) *ErrorValidation {
	var errorFields []ErrorField

	validate := utils.GetDefaultValidator()
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errorField ErrorField
			errorField.Field = toPointer(err.Field())
			errorField.Tag = toPointer(err.Tag())
			errorField.Value = err.Param()

			errorFields = append(errorFields, errorField)
		}

		return &ErrorValidation{
			Fields: errorFields,
		}
	}

	return nil
}

func Convert(err error) *ErrorValidation {
	if err != nil {
		return &ErrorValidation{
			Message: err.Error(),
		}
	}

	return nil
}
