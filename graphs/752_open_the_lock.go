package graphs

import (
	"strconv"
)

func OpenLock(deadends []string, target string) int {
	// Step 1: Set up
	start := "0000"
	visited := make(map[string]bool)
	for _, deadend := range deadends {
		visited[deadend] = true
	}

	// Step 2: Find edge case
	if _, exists := visited[start]; exists {
		return -1
	}

	var increment = func(lock string, pos int) string {
		val, _ := strconv.Atoi(string(lock[pos]))
		valStr := strconv.Itoa((val + 1) % 10)
		res := lock[:pos] + valStr + lock[pos+1:]

		return res
	}

	var decrement = func(lock string, pos int) string {
		val, _ := strconv.Atoi(string(lock[pos]))
		valStr := strconv.Itoa((val - 1 + 10) % 10)
		res := lock[:pos] + valStr + lock[pos+1:]

		return res
	}

	// Step 3: BFS
	queue := []string{start}
	steps := 0

	for len(queue) != 0 {
		elements := queue[:]
		queue = []string{}

		for _, element := range elements {
			if element == target {
				return steps
			}

			for i := range 4 {
				incremented := increment(element, i)
				decremented := decrement(element, i)

				if _, exists := visited[incremented]; !exists {
					queue = append(queue, incremented)
					visited[incremented] = true
				}

				if _, exists := visited[decremented]; !exists {
					queue = append(queue, decremented)
					visited[decremented] = true
				}
			}
		}

		steps++
	}

	return -1
}
