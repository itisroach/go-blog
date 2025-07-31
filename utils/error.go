package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatError(fieldError validator.FieldError) string {
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


type CustomError struct {
	Code 		int
	Message		string
}

func (c *CustomError) Error() *CustomError {
	return c
}