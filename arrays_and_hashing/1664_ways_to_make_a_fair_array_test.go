package arraysandhashing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaysToMakeFair(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name: "ex 1",
			nums: []int{
				2, 1, 6, 4,
			},
			expected: 1,
		},
		{
			name: "ex 2",
			nums: []int{
				1, 1, 1,
			},
			expected: 3,
		},
		{
			name: "ex 3",
			nums: []int{
				1, 2, 3,
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		actual := waysToMakeFair(test.nums)
		if !assert.Equal(t, test.expected, actual) {
			fmt.Printf("fail test %s", test.name)
		}
	}
}
