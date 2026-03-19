package twoddp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name: "example 1",
			prices: []int{
				1, 2, 3, 0, 2,
			},
			expected: 3,
		},
		{
			name: "example 2",
			prices: []int{
				1,
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := maxProfit(test.prices)
			if !assert.Equal(t, test.expected, actual) {
				t.Fatal()
			}
		})
	}
}
