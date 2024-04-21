package leetcode0019

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := []struct {
		name    string
		input   []int
		n       int
		wantRes []int
	}{
		{
			name:    "head =[1,2,3,4,5] n=2",
			input:   []int{1, 2, 3, 4, 5},
			n:       2,
			wantRes: []int{1, 2, 3, 5},
		},
		{
			name:    "head =[1] n=1",
			input:   []int{1},
			n:       1,
			wantRes: []int{},
		},
		{
			name:    "head =[1,2] n=1",
			input:   []int{1, 2},
			n:       1,
			wantRes: []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := removeNthFromEnd(arrayToListNodes(tc.input), tc.n)
			assert.Equal(t, tc.wantRes, listNodesToArray(res))

		})
	}
}

func arrayToListNodes(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	// 新建头指针
	head := &ListNode{Val: nums[0]}
	// 新建 cur 暂存节点用来遍历
	cur := head

	// 遍历数组 1-len(nums)-1
	for i := 1; i < len(nums); i++ {
		newNode := &ListNode{Val: nums[i]}
		// 当前指针的指针域变更
		cur.Next = newNode
		// 当前指针指向新节点
		cur = newNode

	}
	// 返回头指针
	return head
}

func listNodesToArray(head *ListNode) []int {
	cur := head
	res := []int{}
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}
