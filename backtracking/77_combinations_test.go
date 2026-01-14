package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		name     string
		n, k     int
		expected [][]int
	}{
		{
			name:     "example 1",
			n:        4,
			k:        2,
			expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name:     "example 2",
			n:        1,
			k:        1,
			expected: [][]int{{1}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := combine(test.n, test.k)
			if !assert.ElementsMatch(t, output, test.expected) {
				t.Fatalf("%v - expected %v, got %v", test.name, test.expected, output)
			}
		})
	}
}
