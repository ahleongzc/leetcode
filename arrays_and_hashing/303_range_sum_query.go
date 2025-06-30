package arraysandhashing

type NumArray struct {
	nums   []int
	prefix []int
}

func Constructor(nums []int) NumArray {
	numsCopy := make([]int, len(nums))
	prefix := make([]int, len(nums))
	curr := 0
	for i, num := range nums {
		prefix[i] = curr + num
		curr += num
	}

	return NumArray{
		nums:   numsCopy,
		prefix: prefix,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	var leftSum int
	if left != 0 {
		leftSum = this.prefix[left-1]
	}
	return this.prefix[right] - leftSum
}
