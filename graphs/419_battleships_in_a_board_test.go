package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBattleships(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		expected int
	}{
		{
			name: "example 1",
			board: [][]byte{
				{'X', '.', '.', 'X'},
				{'.', '.', '.', 'X'},
				{'.', '.', '.', 'X'},
			},
			expected: 2,
		},
		{
			name: "example 2",
			board: [][]byte{
				{'.'},
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := countBattleships(test.board)
			if !assert.Equal(t, output, test.expected) {
				t.Fatalf("test case: %v - expected: %v, got %v", test.name, test.expected, output)
			}
		})
	}
}
