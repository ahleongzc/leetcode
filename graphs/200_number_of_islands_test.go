package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name: "example 1",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "example 2",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := numIslands(test.grid)
			if !assert.Equal(t, output, test.expected) {
				t.Fatalf(
					"test case: %v - expected: %v, got :%v",
					test.name,
					test.expected,
					output,
				)
			}
		})
	}
}
