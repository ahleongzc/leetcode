package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumConnectedComponents(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		expected int
	}{
		{
			name: "example 1",
			n:    3,
			edges: [][]int{
				{0, 1},
				{0, 2},
			},
			expected: 1,
		},
		{
			name: "example 2",
			n:    6,
			edges: [][]int{
				{0, 1},
				{1, 2},
				{2, 3},
				{4, 5},
			},
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := countComponentsDFS(test.n, test.edges)
			if !assert.Equal(t, output, test.expected) {
				t.Fatalf(
					"test case: %v - expected: %v, got: %v",
					test.name,
					test.expected,
					output,
				)
			}
		})
	}
}
