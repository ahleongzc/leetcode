package slidingwindow

func findClosestElements(arr []int, k int, x int) []int {
	start, end := 0, len(arr)-1

	for end-start > k-1 {
		endDiff := abs(arr[end] - x)
		startDiff := abs(arr[start] - x)

		if startDiff > endDiff {
			start++
		} else {
			end--
		}
	}

	output := make([]int, 0)
	for i := start; i <= end; i++ {
		output = append(output, arr[i])
	}

	return output
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
