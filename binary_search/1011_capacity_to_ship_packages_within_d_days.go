package binarysearch

func ShipWithinDays(weights []int, days int) int {
	left := 1
	right := 0

	for _, weight := range weights {
		right += weight
	}

	canFit := func(capacity int) bool {
		curr := 0
		dayCount := 1

		for _, weight := range weights {
			if weight > capacity {
				return false
			}
			if curr+weight > capacity {
				curr = weight
				dayCount++
			} else {
				curr += weight
			}
		}

		return dayCount <= days
	}

	res := right

	for left <= right {
		mid := (right + left) / 2

		if canFit(mid) {
			right = mid - 1
			res = min(res, mid)
		} else {
			left = mid + 1
		}
	}

	return res
}
