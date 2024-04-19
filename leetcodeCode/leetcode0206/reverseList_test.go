package leetcode0206

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestReverseList(t *testing.T) {
	testCase := []struct {
		name    string
		input   []int
		wantVal []int
	}{
		{
			name:    "输入[1,2,3,4,5]",
			input:   []int{1, 2, 3, 4, 5},
			wantVal: []int{5, 4, 3, 2, 1},
		},
		{
			name:    "输入[1,2]",
			input:   []int{1, 2},
			wantVal: []int{2, 1},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			input := arrayToListNode(tc.input)
			val := reverseList(input)
			assert.Equal(t, arrayToListNode(tc.wantVal), val)
		})
	}
}

// arrayToNode 数组切换为指针
func arrayToListNode(nums []int) *ListNode {

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
