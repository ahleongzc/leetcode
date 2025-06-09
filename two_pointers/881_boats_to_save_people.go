package twopointers

import "slices"

func NumRescueBoats(people []int, limit int) int {
	if len(people) == 0 {
		return 0
	}

	slices.Sort(people)

	count := 0
	left := 0
	right := len(people) - 1

	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}
		right--

		count++
	}

	return count
}
