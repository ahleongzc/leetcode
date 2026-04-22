package intervals

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)

	for i, interval := range intervals {
		// non overlap
		if interval[1] < newInterval[0] {
			res = append(res, interval)
			continue
		}

		if newInterval[1] < interval[0] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			return res
		}

		// overlap
		newMin := min(newInterval[0], interval[0])
		newMax := max(newInterval[1], interval[1])
		newInterval = []int{newMin, newMax}
	}

	res = append(res, newInterval)
	return res
}
