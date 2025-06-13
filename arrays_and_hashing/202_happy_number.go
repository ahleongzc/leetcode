package arraysandhashing

import "strconv"

func IsHappy(n int) bool {
	appeared := make(map[string]bool)

	for {
		numberInString := strconv.Itoa(n)
		if _, exists := appeared[numberInString]; exists {
			return false
		}
		appeared[numberInString] = true

		sum := 0
		for _, number := range numberInString {
			numberInInt := int(number - '0')
			sum += numberInInt * numberInInt
		}

		if sum == 1 {
			return true
		}
		n = sum
	}
}
