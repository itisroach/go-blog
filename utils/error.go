package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func formatError(fieldError validator.FieldError) string {
	field := strings.ToLower(fieldError.Field())

	switch fieldError.Tag() {
	case "required":
		return field + " is required"
	case "min":
		return field + " must be at least " + fieldError.Param() + " characters"
	case "max":
		return field + " must be at most " + fieldError.Param() + " characters"
	}

	return field + " is invalid"
}


func GenerateUserFriendlyError(err error) *[]string {
	// using field error validator from gorm
	var validator validator.ValidationErrors

	// convert the error to a user friendly error
	if errors.As(err, &validator) {
		out := make([]string, len(validator))

		for i, fieldError := range validator {
			out[i] = formatError(fieldError)
		}
			
		
		return &out
	}

	return nil
}

type CustomError struct {
	Code 		int
	Message		string
}

func (c *CustomError) Error() *CustomError {
	return c
}