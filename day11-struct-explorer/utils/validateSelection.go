package utils

import "slices"

func ValidateSelection(str string) bool {
	allowedOpens := []string{"1", "2", "3", "4"}
	return slices.Contains(allowedOpens, str)
}
