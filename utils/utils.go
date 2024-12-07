package utils

import (
	"fmt"
	"math/big"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func FormatMoney(amount *big.Float) string {
	return amount.Text('f', 2)
}

func ParseMoney(amount string) (*big.Float, error) {
	f, ok := new(big.Float).SetString(amount)
	if !ok {
		return nil, fmt.Errorf("invalid amount")
	}
	return f, nil
}

func GetDefaultValidator() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		tag := field.Tag.Get("json")
		if tag == "-" {
			return ""
		}

		return tag
	})

	return validate
}

func FormatDate(date time.Time) string {
	return date.Format(time.DateOnly)
}
