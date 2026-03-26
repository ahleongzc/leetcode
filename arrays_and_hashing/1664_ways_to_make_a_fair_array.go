package arraysandhashing

func waysToMakeFair(nums []int) int {
	prefixEven := make([]int, len(nums)+2)
	prefixOdd := make([]int, len(nums)+2)

	for i := range nums {
		indexPrefixArr := i + 1
		prefixEven[indexPrefixArr] = prefixEven[indexPrefixArr-1]
		prefixOdd[indexPrefixArr] = prefixOdd[indexPrefixArr-1]
		if i%2 == 0 {
			prefixEven[indexPrefixArr] += nums[i]
		} else {
			prefixOdd[indexPrefixArr] += nums[i]
		}
	}

	reversedPrefixEven := make([]int, len(nums)+2)
	reversedPrefixOdd := make([]int, len(nums)+2)
	for i := len(nums) - 1; i > -1; i-- {
		indexPrefixArr := i + 1
		reversedPrefixEven[indexPrefixArr] = reversedPrefixEven[indexPrefixArr+1]
		reversedPrefixOdd[indexPrefixArr] = reversedPrefixOdd[indexPrefixArr+1]
		if i%2 == 0 {
			reversedPrefixEven[indexPrefixArr] += nums[i]
		} else {
			reversedPrefixOdd[indexPrefixArr] += nums[i]
		}
	}

	count := 0
	for i := 1; i < len(reversedPrefixEven)-1; i++ {
		if prefixEven[i-1]+reversedPrefixOdd[i+1] == prefixOdd[i-1]+reversedPrefixEven[i+1] {
			count++
		}
	}
	return count
}
