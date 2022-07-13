package customValidators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateCoolTitles(filds validator.FieldLevel) bool {
	return strings.Contains(filds.Field().String(), "cool")
}
