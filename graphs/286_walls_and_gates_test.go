package graphs

import (
	"reflect"
	"testing"
)

func TestWallsAndGates(t *testing.T) {
	const INF = 2147483647

	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name: "basic case",
			input: [][]int{
				{INF, -1, 0, INF},
				{INF, INF, INF, -1},
				{INF, -1, INF, -1},
				{0, -1, INF, INF},
			},
			expected: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			name: "no gates",
			input: [][]int{
				{INF, INF},
				{INF, INF},
			},
			expected: [][]int{
				{INF, INF},
				{INF, INF},
			},
		},
		{
			name: "all gates",
			input: [][]int{
				{0, 0},
				{0, 0},
			},
			expected: [][]int{
				{0, 0},
				{0, 0},
			},
		},
		{
			name: "single gate with walls",
			input: [][]int{
				{INF, -1, 0},
				{INF, -1, INF},
			},
			expected: [][]int{
				{INF, -1, 0},
				{INF, -1, 1},
			},
		},
		{
			name: "leetcode example 1",
			input: [][]int{
				{INF, -1, 0, INF},
				{INF, INF, INF, -1},
				{INF, -1, INF, -1},
				{0, -1, INF, INF},
			},
			expected: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			name: "leetcode example 2",
			input: [][]int{
				{0, -1},
				{INF, INF},
			},
			expected: [][]int{
				{0, -1},
				{1, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WallsAndGates(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("got %v, want %v", tt.input, tt.expected)
			}
		})
	}
}
