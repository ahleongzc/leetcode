package binarysearch

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{
			nums: []int{3, 4, 5, 1, 2},
			ans:  1,
		},
		{
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			ans:  0,
		},
		{
			nums: []int{11, 13, 15, 17},
			ans:  11,
		},
	}

	for i, tt := range tests {
		t.Run(
			fmt.Sprintf("test %d", i),
			func(t *testing.T) {
				res := FindMin(tt.nums)
				if res != tt.ans {
					t.Fatalf("expected: %d, returned: %d", tt.ans, res)
				}
			},
		)
	}
}
