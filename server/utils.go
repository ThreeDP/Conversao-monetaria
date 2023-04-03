package server

import (
	"unicode"
)

func IsAllDigit(value string) bool {
	dot := false
	if len(value) <= 0 {
		return false
	}
	for i := 0; i < len(value); i++ {
		if rune(value[i]) == '.' && !dot {
			dot = true
			continue
		}
		if !unicode.IsDigit(rune(value[i])) {
			return false
		}
	}
	return true
}