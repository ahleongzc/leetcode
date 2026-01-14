package arraysandhashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		target   int
		expected []int
	}{
		{
			name:     "example 1",
			numbers:  []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			name:     "example 2",
			numbers:  []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		{
			name:     "example 3",
			numbers:  []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		result := twoSumII(test.numbers, test.target)
		if !assert.ElementsMatch(t, result, test.expected) {
			t.Fatalf("expected %v, got %v", test.expected, result)
		}
	}
}
