package twopointers

func ValidPalindromeTwoO1SpaceComplexity(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if s[start] != s[end] {
			return ValidPalindromTwoO1SpaceComplexityHelper(s, start+1, end) ||
				ValidPalindromTwoO1SpaceComplexityHelper(s, start, end-1)
		}
		start++
		end--
	}

	return true
}

func ValidPalindromTwoO1SpaceComplexityHelper(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func ValidPalindromeTwo(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if s[start] != s[end] {
			removeStart := s[:start] + s[start+1:]
			removeEnd := s[:end] + s[end+1:]
			return IsValidPalindromeTwoHelper(removeStart) || IsValidPalindromeTwoHelper(removeEnd)
		}
		start++
		end--
	}

	return true
}

func IsValidPalindromeTwoHelper(s string) bool {
	reversedRunes := make([]rune, 0)

	for i := len(s) - 1; i >= 0; i-- {
		reversedRunes = append(reversedRunes, rune(s[i]))
	}

	originalRunes := []rune(s)

	for i, r := range originalRunes {
		if reversedRunes[i] != r {
			return false
		}
	}

	return true
}
