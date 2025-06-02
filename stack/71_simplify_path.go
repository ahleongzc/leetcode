package stack

import (
	"strings"
)

func SimplifyPath(path string) string {
	if len(path) == 1 {
		return path
	}

	splittedPath := strings.Split(path, "/")
	stack := make([]string, 0)

	for _, sp := range splittedPath {
		if sp == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}

		if sp == "." || sp == "" {
			continue
		}

		stack = append(stack, sp)
	}

	return "/" + strings.Join(stack, "/")
}
