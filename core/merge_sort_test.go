package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:  "one element",
			input: []int{1},
		},
		{
			name:  "5 elements",
			input: []int{4, 3, 5, 1, 2},
		},
		{
			name:  "5 elements with negative numbers",
			input: []int{-4, 3, 5, -1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := mergeSort(test.input)
			if len(output) != len(test.input) {
				t.Fatalf("test case: %v - the length of the output must be the same", test.name)
			}

			if !assert.IsIncreasing(t, output) {
				t.Fatalf("test case: %v - the output is not sorted", test.name)
			}
		})
	}
}
