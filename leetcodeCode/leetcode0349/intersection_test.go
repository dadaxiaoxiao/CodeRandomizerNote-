package leetcode0349

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestIntersection(t *testing.T) {
	testCases := []struct {
		name    string
		nums1   []int
		nums2   []int
		wantRes []int
	}{
		{
			name:    "nums1=[1,2,2,1],nums2 =[2,2]",
			nums1:   []int{1, 2, 2, 1},
			nums2:   []int{2, 2},
			wantRes: []int{2},
		},
		{
			name:    "nums1=[4,9,5],nums2 =[9,4,9,8,4]",
			nums1:   []int{4, 9, 5},
			nums2:   []int{9, 4, 9, 8, 4},
			wantRes: []int{9, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := intersection(tc.nums1, tc.nums2)
			assert.Equal(t, tc.wantRes, res)

		})
	}
}
