package stack

import "strconv"

func CalPoints(operations []string) int {
	points := 0
	stack := []int{}
	prevOne := 0
	prevTwo := 0

	for _, op := range operations {
		if op == "+" {
			prevOne = stack[len(stack) - 1]
			prevTwo = stack[len(stack) - 2]
			stack = append(stack, prevOne + prevTwo)
			continue
		}

		if op == "C" {
			stack = stack[:len(stack) - 1]
			continue
		}

		if op == "D" {
			prevOne = stack[len(stack) - 1]
			stack = append(stack, prevOne * 2)
			continue
		}
	
		val, _ := strconv.Atoi(op)
		stack = append(stack, val)
	}

	for _, val := range stack {
		points += val
	}
	
	return points
}
