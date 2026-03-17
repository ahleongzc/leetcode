package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			name: "example 1",
			matrix: [][]int{
				{1, 2, 4, 8},
				{10, 11, 12, 13},
				{14, 20, 30, 40},
			},
			target:   10,
			expected: true,
		},
		{
			name: "example 2",
			matrix: [][]int{
				{1, 2, 4, 8},
				{10, 11, 12, 13},
				{14, 20, 30, 40},
			},
			target:   15,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := searchMatrix(test.matrix, test.target)
			if !assert.Equal(t, test.expected, actual) {
				t.Fatal()
			}
		})
	}

}
