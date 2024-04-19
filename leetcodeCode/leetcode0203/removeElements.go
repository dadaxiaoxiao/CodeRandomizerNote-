package leetcode0203

// ListNode 指针
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeElements 移除链表元素
// leetcode203
func removeElements(head *ListNode, val int) *ListNode {
	// 新建虚拟头节点
	dummyNode := &ListNode{}
	// 为了避免删除头节点的情况，这里使用统一的删除规则
	dummyNode.Next = head
	// 新建 cur 暂存节点用来遍历
	cur := dummyNode

	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			// 变更当前指针的指针域
			cur.Next = cur.Next.Next
		} else {
			// 当前指针指向下一个
			cur = cur.Next
		}
	}
	return dummyNode.Next
}

// arrayToNode 数组切换为指针
func arrayToNode(nums []int) *ListNode {

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
