package arraysandhashing

func ProductExceptSelfO1Space(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i := range nums {
		res[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}

func ProductExceptSelf(nums []int) []int {
	productLeft := make([]int, len(nums))
	productLeft[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		productLeft[i] = nums[i] * productLeft[i-1]
	}

	productRight := make([]int, len(nums))
	productRight[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		productRight[i] = nums[i] * productRight[i+1]
	}

	res := make([]int, len(nums))
	res[0] = productRight[1]
	res[len(nums)-1] = productLeft[len(nums)-2]

	for i := 1; i < len(nums)-1; i++ {
		res[i] = productLeft[i-1] * productRight[i+1]
	}

	return res
}
