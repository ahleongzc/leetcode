package slidingwindow

// A B B B B
// L
//     R

func CharacterReplacement(s string, k int) int {
	countMap := make(map[byte]int)
	left, right := 0, 0
	count := 0

	for right < len(s) {
		countMap[s[right]]++

		if (right-left+1)-findLargest(countMap) > k {
			countMap[s[left]]--
			left++
		}

		count = max(right-left+1, count)
		right++
	}

	return count
}

func findLargest(countMap map[byte]int) int {
	res := 0
	for _, v := range countMap {
		res = max(res, v)
	}

	return res
}
