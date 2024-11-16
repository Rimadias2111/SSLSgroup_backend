package filters

import (
	"regexp"
)

func ValidatePhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)

	if !re.MatchString(phone) {
		return false
	}
	return true
}
