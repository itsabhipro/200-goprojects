package utils

import (
	"regexp"
)

func ValidateEmail(email string) bool {
	if email[0] == '.' {
		return false
	}
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

	return re.Match([]byte(email))
}
