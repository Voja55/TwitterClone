package validation

import (
	"strings"
)

func ValidateFirstName(fn string) bool {
	for _, c := range fn {
		if strings.ContainsRune("~!@#$%^&*()_+{}:|';<>?", c) {
			return false
		}
		if strings.ContainsRune("1234567890", c) {
			return false
		}
	}
	return true
}

func ValidateLastName(fn string) bool {
	for _, c := range fn {
		if strings.ContainsRune("~!@#$%^&*()_+{}:|';<>?", c) {
			return false
		}
		if strings.ContainsRune("1234567890", c) {
			return false
		}
	}
	return true
}

func ValidateAddress(fn string) bool {
	for _, c := range fn {
		if strings.ContainsRune("~!@#$%^&*()_+{}:|';<>?", c) {
			return false
		}
	}
	return true
}

func ValidateCompanyName(fn string) bool {
	for _, c := range fn {
		if strings.ContainsRune("~!@#$%^&*()_+{}:|';<>?", c) {
			return false
		}
	}
	return true
}

func ValidateWebSite(fn string) bool {
	for _, c := range fn {
		if strings.ContainsRune("~!@#$%^&*()_+{}:|';<>?", c) {
			return false
		}
		if strings.ContainsRune("1234567890", c) {
			return false
		}
	}
	return true
}
