package leetcode0977

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSortedSquares(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		wantNums []int
	}{
		{
			name:     "输入 [-4, -1, 0, 3, 10]",
			nums:     []int{-4, -1, 0, 3, 10},
			wantNums: []int{0, 1, 9, 16, 100},
		},
		{
			name:     "输入 [-7,-3,2,3,11]",
			nums:     []int{-7, -3, 2, 3, 11},
			wantNums: []int{4, 9, 9, 49, 121},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := sortedSquares(tc.nums)
			assert.Equal(t, tc.wantNums, res)
		})
	}
}
