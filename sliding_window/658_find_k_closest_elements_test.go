package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		k        int
		x        int
		expected []int
	}{
		{
			name: "example 1",
			arr:  []int{1, 2, 3, 4, 5},
			k:    4,
			x:    3,
			expected: []int{
				1, 2, 3, 4,
			},
		},
		{
			name: "example 2",
			arr:  []int{1, 1, 2, 3, 4, 5},
			k:    4,
			x:    -1,
			expected: []int{
				1, 1, 2, 3,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findClosestElements(test.arr, test.k, test.x)
			if !assert.ElementsMatch(t, result, test.expected) {
				t.Fatalf("%v: expected %v, got %v", test.name, test.expected, result)
			}
		})
	}
}
