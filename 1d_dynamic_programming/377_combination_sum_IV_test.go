package oneddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum4(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name: "example 1",
			nums: []int{
				1, 2, 3,
			},
			target:   4,
			expected: 7,
		},
		{
			name: "example 2",
			nums: []int{
				9,
			},
			target:   3,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := combinationSum4(test.nums, test.target)
			if !assert.Equal(t, test.expected, actual) {
				t.Fatal()
			}
		})
	}
}
