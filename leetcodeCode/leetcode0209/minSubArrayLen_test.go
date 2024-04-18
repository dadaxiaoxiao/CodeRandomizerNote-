package leetcode0209

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMinSubArrayLen(t *testing.T) {
	testCases := []struct {
		name    string
		nums    []int
		target  int
		wantRes int
	}{
		{
			name:    "target =7,nums =[2,3,1,2,4,3],预期 2",
			nums:    []int{2, 3, 1, 2, 4, 3},
			target:  7,
			wantRes: 2,
		},
		{
			name:    "target =4,nums =[1,4,4],预期 1",
			nums:    []int{1, 4, 4},
			target:  4,
			wantRes: 1,
		},
		{
			name:    "target =11,nums =[1,1,1,1,1,1,1,1],预期 0",
			nums:    []int{1, 1, 1, 1, 1, 1, 1, 1},
			target:  11,
			wantRes: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := minSubArrayLen(tc.target, tc.nums)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
