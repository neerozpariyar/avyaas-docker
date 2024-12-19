package utils

import (
	"net/mail"
	"regexp"
)

// IsValidEmail checks if the given string data is a valid email string.
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	// Check if there's a dot after the @ symbol
	match, _ := regexp.MatchString(`@.+\..+`, email)
	return match
}
