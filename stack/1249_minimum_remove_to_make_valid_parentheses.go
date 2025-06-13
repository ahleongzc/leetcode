package stack

import (
	"slices"
	"unicode"
)

func MinRemoveToMakeValid(s string) string {
	validWord := s

	type element struct {
		char  rune
		index int
	}

	stack := make([]element, 0)

	for i, char := range s {
		if unicode.IsLetter(char) {
			continue
		}

		if char == '(' {
			stack = append(stack, element{
				char:  char,
				index: i,
			})
		}

		if char == ')' {
			if len(stack) == 0 || stack[len(stack)-1].char != rune('(') {
				stack = append(stack, element{
					char:  char,
					index: i,
				})
				continue
			}
			stack = stack[:len(stack)-1]
		}
	}

	slices.Reverse(stack)

	for _, element := range stack {
		validWord = validWord[:element.index] + validWord[element.index+1:]
	}

	return validWord
}
