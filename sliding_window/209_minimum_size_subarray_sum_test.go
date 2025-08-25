package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	target int
	nums   []int
	res    int
}{
	{
		target: 7,
		nums:   []int{2, 3, 1, 2, 4, 3},
		res:    2,
	},
	{
		target: 4,
		nums:   []int{1, 4, 4},
		res:    1,
	},
	{
		target: 11,
		nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
		res:    0,
	},
	{
		target: 4,
		nums:   []int{1, 4, 4},
		res:    1,
	},
}

func TestMinSubArrayLen(t *testing.T) {
	for _, tt := range testCases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tt.res, MinSubArrayLen(tt.target, tt.nums))
		})
	}
}
