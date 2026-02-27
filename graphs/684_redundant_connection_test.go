package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRedundantConnections(t *testing.T) {
	tests := []struct {
		name     string
		edges    [][]int
		expected []int
	}{
		{
			name: "example 1",
			edges: [][]int{
				{1, 2}, {1, 3}, {2, 3},
			},
			expected: []int{
				2, 3,
			},
		},
		{
			name: "example 2",
			edges: [][]int{
				{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5},
			},
			expected: []int{
				1, 4,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := findRedundantConnection(test.edges)

			if !assert.ElementsMatch(t, output, test.expected) {
				t.Fatalf(
					"test %v: expected %v, got %v",
					test.name,
					test.expected,
					output,
				)
			}
		})
	}
}
