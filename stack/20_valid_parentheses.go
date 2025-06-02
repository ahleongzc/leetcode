package stack

func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0)

	for _, char := range s {
		if char == '(' {
			stack = append(stack, ')')
			continue
		}

		if char == '[' {
			stack = append(stack, ']')
			continue
		}

		if char == '{' {
			stack = append(stack, '}')
			continue
		}

		if len(stack) == 0 {
			return false
		}

		lastChar := stack[len(stack)-1]
		if lastChar != char {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
