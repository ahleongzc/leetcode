package greedy

import (
	"slices"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
	slices.Sort(asteroids)
	curr := mass

	for _, asteroid := range asteroids {
		if curr < asteroid {
			return false
		}
		curr += asteroid
	}

	return true
}
