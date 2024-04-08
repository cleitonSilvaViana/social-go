package fail

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

// FieldError ...
type FieldError struct {
	Field   string `json:"fild"`
	Message string `json:"message"`
}

func getFieldError(field validator.FieldError) string {
	switch field.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + field.Param()
	case "gte":
		return "Should be greater than " + field.Param()
	case "gt":
		return "should be greater than " + field.Param()
	case "email":
		return "Your email is invalid"
	}
	return "Unknown error"
}

func ValidateFields(err error) *validator.ValidationErrors {
	var fieldError validator.ValidationErrors

	if errors.As(err, &fieldError) {
		output := make([]FieldError, len(fieldError))
		for i, fe := range fieldError {
			output[i] = FieldError{fe.Field(), getFieldError(fe)}
		}
		return &fieldError
	}
	return nil
}

