package leetcode0704

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name string
		// 输入
		nums   []int
		target int
		// 输出
		wantIndex int
	}{
		{
			name:      "9 出现在 nums 中并且下标为 4",
			nums:      []int{-1, 0, 3, 5, 9, 12},
			target:    9,
			wantIndex: 4,
		},
		{
			name:      "2 不存在 nums 中因此返回 -1",
			nums:      []int{-1, 0, 3, 5, 9, 12},
			target:    2,
			wantIndex: -1,
		},
		{
			name:      "13 不存在 nums 中因此返回 -1",
			nums:      []int{-1, 0, 3, 5, 9, 12},
			target:    13,
			wantIndex: -1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := search2(tc.nums, tc.target)
			assert.Equal(t, tc.wantIndex, res)
		})
	}
}
