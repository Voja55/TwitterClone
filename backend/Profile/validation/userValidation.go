package validation

import (
	"strings"
)

const illegalChars = "~!@#$%^&*()_+{}:|';<>?"

const numbers = "1234567890"

func ValidateCharAndNum(fn string) bool {
	for _, c := range fn {
		if strings.ContainsRune(illegalChars, c) {
			return false
		}
		if strings.ContainsRune(numbers, c) {
			return false
		}
	}
	return true
}

func ValidateChar(fn string) bool {
	for _, c := range fn {
		if strings.ContainsRune(illegalChars, c) {
			return false
		}
	}
	return true
}
