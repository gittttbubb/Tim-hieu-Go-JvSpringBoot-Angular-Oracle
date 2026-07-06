package validator

import (
	"errors"
	"regexp"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("Password must be at least 8 characters")
	}

	if len(password) > 100 {
		return errors.New("Password must not exceed 100 characters")
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	if !hasUpper {
		return errors.New("Password must contain at least one uppercase letter")
	}

	if !hasLower {
		return errors.New("Password must contain at least one lowercase letter")
	}

	if !hasNumber {
		return errors.New("Password must contain at least one number")
	}

	if !hasSpecial {
		return errors.New("Password must contain at least one special character")
	}

	return nil
}