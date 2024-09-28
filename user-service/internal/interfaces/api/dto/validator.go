package dto

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetValidator() *validator.Validate {
	return validate
}

func init() {
	// Register the custom validation function
	_ = GetValidator().RegisterValidation("iranphone", iranPhoneValidator)
}

// Define your regex pattern for Iranian phone numbers (E.164 format with exactly 10 digits after +98).
var iranPhoneRegex = regexp.MustCompile(`^\+98\d{10}$`)

// Custom validation function for phone numbers
func iranPhoneValidator(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	return iranPhoneRegex.MatchString(phone)
}
