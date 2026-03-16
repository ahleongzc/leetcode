package oneddp

func maxProduct(nums []int) int {
	res := nums[0]
	currMin, currMax := 1, 1

	for _, num := range nums {
		tmp := currMax * num

		// if the current value is 0, then the next iteration will just take the number itself,
		// updating currMax to be a positive value (..., num)
		currMax = max(tmp, num*currMin, num)
		currMin = min(tmp, num*currMin, num)
		res = max(res, currMax)
	}

	return res
}
