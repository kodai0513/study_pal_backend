package validations

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	errMessageValidationRequired     = "this field is required"
	errMessageValidationAlphanumeric = "this field must contain only letters and numbers"
	errMessageValidationEmail        = "invalid email format"
	errMessageValidationMinLength    = "this field must be at least %d characters long"
	errMessageValidationMaxLength    = "this field must be at most %d characters long"
)

func IsAlphanumeric(value string) (bool, error) {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)

	if !re.MatchString(value) {
		return false, errors.New(errMessageValidationAlphanumeric)
	}

	return true, nil
}

func IsEmailValid(value string) (bool, error) {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9]+(?:\.[a-zA-Z0-9-]+)+$`)
	if !emailRegex.MatchString(value) {
		return false, errors.New(errMessageValidationEmail)
	}

	return true, nil
}

func IsCharactersMaxLength(value string, maxLength uint) (bool, error) {
	if len(value) > int(maxLength) && len(value) > 0 {
		return false, fmt.Errorf(errMessageValidationMaxLength, maxLength)
	}

	return true, nil
}

func IsCharactersMinLength(value string, minLength uint) (bool, error) {
	if len(value) < int(minLength) && len(value) > 0 {
		return false, fmt.Errorf(errMessageValidationMinLength, minLength)
	}

	return true, nil
}

func IsRequired(value string) (bool, error) {
	if len(value) <= 0 {
		return false, errors.New(errMessageValidationRequired)
	}

	return true, nil
}
