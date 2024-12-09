package validations

import (
	"credit/utils"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
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
	SanitizeStruct(&payload)

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

func SanitizeStruct[T any](payload *T) {
	p := bluemonday.UGCPolicy()

	val := reflect.ValueOf(payload).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String && field.CanSet() {
			originalValue := field.String()
			sanitizedValue := p.Sanitize(originalValue)
			field.SetString(sanitizedValue)
		}

		if field.Kind() == reflect.Struct {
			nestedStruct := field.Addr().Interface()
			SanitizeStruct(&nestedStruct)
		}
	}
}

func Convert(err error) *ErrorValidation {
	if err != nil {
		return &ErrorValidation{
			Message: err.Error(),
		}
	}

	return nil
}
