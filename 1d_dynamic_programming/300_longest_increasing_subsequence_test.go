package oneddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{9, 1, 4, 2, 3, 3, 7},
			expected: 4,
		},
		{
			name:     "example 2",
			nums:     []int{0, 3, 1, 3, 2, 3},
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := lengthOfLIS(test.nums)
			if !assert.Equal(t, test.expected, actual) {
				t.Fatal()
			}
		})
	}
}
