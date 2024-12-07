package utils_tests

import (
	"credit/utils"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUtils_HashAndCheckPassword(t *testing.T) {
	password := "password123"

	hashed, err := utils.HashPassword(password)
	assert.NoError(t, err)

	valid := utils.CheckPasswordHash(password, hashed)
	assert.Equal(t, valid, true)
}

func TestUtils_FormatAndParseMoney(t *testing.T) {
	money := new(big.Float).SetFloat64(500000.25)

	formatted := utils.FormatMoney(money)

	parsedMoney, err := utils.ParseMoney(formatted)

	assert.NoError(t, err)
	assert.True(t, money.Cmp(parsedMoney) == 0)
}

func TestGetDefaultValidator(t *testing.T) {
	validate := utils.GetDefaultValidator()

	if validate == nil {
		t.Error("Expected validator instance to be non-nil, got nil")
	}

	type TestStruct struct {
		Field1 string `json:"field_1"`
		Field2 string `json:"-"`
	}

	field, ok := reflect.TypeOf(TestStruct{}).FieldByName("Field1")
	if !ok {
		t.Fatal("Field 'Field1' not found in TestStruct")
	}

	tagName := getTagName(field)

	expectedTagName := "field_1"
	if tagName != expectedTagName {
		t.Errorf("Expected tag name '%s', got '%s'", expectedTagName, tagName)
	}

	field, ok = reflect.TypeOf(TestStruct{}).FieldByName("Field2")
	if !ok {
		t.Fatal("Field 'Field2' not found in TestStruct")
	}

	tagName = getTagName(field)

	expectedTagName = ""
	if tagName != expectedTagName {
		t.Errorf("Expected tag name '%s', got '%s'", expectedTagName, tagName)
	}
}

func TestFormatDate(t *testing.T) {
	input := time.Date(1990, time.February, 1, 0, 0, 0, 0, time.UTC)
	expected := "1990-02-01"
	result := utils.FormatDate(input)

	assert.Equal(t, expected, result)
}

func getTagName(field reflect.StructField) string {
	tag := field.Tag.Get("json")
	if tag == "-" {
		return ""
	}
	return tag
}
