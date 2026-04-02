package intervals

import "slices"

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	for i := 1; i < len(intervals); i++ {
		prev := intervals[i-1]
		curr := intervals[i]

		if curr[0] <= prev[1] {
			curr[0] = prev[0]
			prev[0] = -1
			curr[1] = max(prev[1], curr[1])
			continue
		}
	}

	for _, interval := range intervals {
		if interval[0] != -1 {
			res = append(res, interval)
		}
	}

	return res
}
