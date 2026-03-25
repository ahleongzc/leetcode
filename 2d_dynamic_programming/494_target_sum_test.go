package twoddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "ex 1",
			nums:     []int{2, 2, 2},
			target:   2,
			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := findTargetSumWays(test.nums, test.target)
			if !assert.Equal(t, test.expected, actual) {
				t.Fatal(test.name)
			}
		})
	}
}
