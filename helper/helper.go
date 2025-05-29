package helper

import (
	"unicode"
)

func IsRuneAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func ReverseString(s string) string {
	runes := make([]rune, 0)

	for i := len(s) - 1; i >= 0; i-- {
		runes = append(runes, rune(s[i]))
	}

	return string(runes)
}

func RemoveRuneFromString(s string, i int) string{
	return s[:i] + s[i+1:]
}
