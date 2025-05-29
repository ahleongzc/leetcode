package twopointers

import (
	"unicode"
)

func IsPalindromeO1SpaceComplexity(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		for !IsRuneAlphanumeric(rune(s[start])) {
			start++
		}
		for !IsRuneAlphanumeric(rune(s[end])) {
			end--
		}

		if start > end {
			break
		}
		if unicode.ToLower(rune(s[start])) != unicode.ToLower(rune(s[end])) {
			return false
		}

		start++
		end--
	}

	return true
}

func IsRuneAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func IsPalindrome(s string) bool {
	runeArr := []rune{}

	for _, r := range s {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			continue
		}
		runeArr = append(runeArr, unicode.ToLower(r))
	}

	start := 0
	end := len(runeArr) - 1

	for start < end {
		if runeArr[start] != runeArr[end] {
			return false
		}

		start++
		end--
	}

	return true
}
