package leetcode00027

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name       string
		nums       []int
		val        int
		wantLength int
	}{
		{
			name:       "输入：nums = [3,2,2,3], val = 3，输出 2",
			nums:       []int{3, 2, 2, 3},
			val:        3,
			wantLength: 2,
		},
		{
			name:       "输入：nums = [0,1,2,2,3,0,4,2], val = 2，输出 5",
			nums:       []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:        2,
			wantLength: 5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := removeElement(tc.nums, tc.val)
			assert.Equal(t, tc.wantLength, res)
		})
	}
}
