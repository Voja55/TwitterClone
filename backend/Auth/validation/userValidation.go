package validation

import (
	"strings"
	"unicode"
)

func ValidatePassword(pw string) bool {
	if len(pw) < 8 {
		return false
	}
	upper := 0
	lower := 0
	special := 0
	number := 0
	for _,c := range pw {
		if unicode.IsUpper(c) {
			upper ++
		}
		if unicode.IsLower(c) {
			lower ++
		}
		if unicode.IsDigit(c) {
			number ++
		}
		if strings.ContainsRune("-,@$!._?&", c) {
			special ++
		}
	}
	return upper > 0 && lower > 0 && special > 0 && number > 0
}

func ValidateUsername(un string) bool {
	if len(un) < 5 || len(un) > 25 {
		return false
	}
	runes := []rune(un)
	previousc := 'a'
	for i,c := range runes {
		if unicode.IsUpper(c) {return false}
		if i == 0 && strings.ContainsRune("._", c) {return false}
		if strings.ContainsRune(".", c) && strings.ContainsRune(".", previousc) {return false}
		if strings.ContainsRune("_", c) && strings.ContainsRune("_", previousc) {return false}
		previousc = c
	}
	if runes[len(runes)-1] == '_' || runes[len(runes)-1] == '.' {return false}
	return true
}

func ValidateRole(r string) bool {
	if r == "regular" || r == "buisiness" {return true} 
	return false
}