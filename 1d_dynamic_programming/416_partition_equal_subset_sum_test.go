package oneddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartition(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name: "ex 1",
			nums: []int{
				1, 5, 11, 5,
			},
			expected: true,
		},
		{
			name: "ex 2",
			nums: []int{
				1, 2, 3, 5,
			},
			expected: false,
		},
	}

	for _, test := range tests {
		actual := canPartition(test.nums)
		if !assert.Equal(t, test.expected, actual) {
			t.Fatal(test.name)
		}
	}
}
