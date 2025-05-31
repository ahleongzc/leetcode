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

	// Step 2: Check if start is in deadends (edge case)
	if _, ok := visited[start]; ok {
		return -1
	}

	// Step 3: Helper functions to increment or decrement
	var increment = func(lock string, pos int) string {
		val, _ := strconv.Atoi(string(lock[pos]))
		posVal := strconv.Itoa((val + 1) % 10)
		res := lock[:pos] + posVal + lock[pos+1:]

		return res
	}

	var decrement = func(lock string, pos int) string {
		val, _ := strconv.Atoi(string(lock[pos]))
		posVal := strconv.Itoa((val - 1 + 10) % 10)
		res := lock[:pos] + posVal + lock[pos+1:]

		return res
	}

	// Step 4: BFS
	queue := make([]string, 0)
	queue = append(queue, start)
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

				if _, ok := visited[incremented]; !ok {
					queue = append(queue, incremented)
					visited[incremented] = true
				}
				if _, ok := visited[decremented]; !ok {
					queue = append(queue, decremented)
					visited[decremented] = true
				}
			}
		}

		steps++
	}

	return -1
}
