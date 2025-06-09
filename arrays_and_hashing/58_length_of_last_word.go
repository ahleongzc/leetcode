package arraysandhashing

import (
	"unicode"
)

func LengthOfLastWord(s string) int {
	ptr := len(s) - 1

	for ptr >= 0 && unicode.IsSpace(rune(s[ptr])) {
		ptr--
	}

	length := 0
	for ptr >= 0 && unicode.IsLetter(rune(s[ptr])) {
		ptr--
		length++
	}

	return length
}
    
