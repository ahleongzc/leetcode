package intervals

import "slices"

func CanAttendMeetings(intervals [][]int) bool {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] > intervals[i][0] {
			return false
		}
	}

	return true
}
