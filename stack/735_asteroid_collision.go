package stack

func AsteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	stack = append(stack, asteroids[0])

	for i := 1; i < len(asteroids); i++ {
		curr := asteroids[i]

		for len(stack) > 0 && curr < 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]
			res := curr + top

			if res == 0 {
				stack = stack[:len(stack)-1]
				curr = 0
			} else if res > 0 {
				curr = 0
			} else {
				stack = stack[:len(stack)-1]
			}
		}

		if curr == 0 {
			continue
		}

		stack = append(stack, curr)
	}
	return stack
}
