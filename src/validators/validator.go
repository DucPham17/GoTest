package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidatorField(fielf validator.FieldLevel) bool {
	return strings.Contains(fielf.Field().String(), "cool")
}
