package leetcode0203

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestRemoveElements(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		val      int
		wantHead *ListNode
	}{
		{
			name:     "输入 [1,2,6,3,4,5,6] val = 6",
			head:     arrayToNode([]int{1, 2, 6, 3, 4, 5, 6}),
			val:      6,
			wantHead: arrayToNode([]int{1, 2, 3, 4, 5}),
		},
		{
			name:     "输入 [] val = 1",
			head:     arrayToNode([]int{}),
			val:      1,
			wantHead: arrayToNode([]int{}),
		},
		{
			name:     "输入 [7,7,7,7] val = 7",
			head:     arrayToNode([]int{7, 7, 7, 7}),
			val:      7,
			wantHead: arrayToNode([]int{}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := removeElements(tc.head, tc.val)
			assert.Equal(t, tc.wantHead, res)
		})
	}
}
