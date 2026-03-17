package binarysearch

import "slices"

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, slices.Max(piles)

	var canFinish func(speed int) bool
	canFinish = func(speed int) bool {
		hoursTaken := 0
		for _, pile := range piles {
			hoursTaken += pile / speed
			if pile%speed != 0 {
				hoursTaken++
			}
		}
		return hoursTaken <= h
	}

	for left < right {
		mid := (left + right) / 2
		if canFinish(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
