package leetcode0024

type ListNode struct {
	Val  int
	Next *ListNode
}

// swapPairs
// leetcode2
// 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}

	// 暂存节点用来遍历
	curr := dummyHead

	// 节点数目为偶数   curr.Next ！= nil
	// 节点数目为奇数   curr.Next.Next != nil
	for curr.Next != nil && curr.Next.Next != nil {
		node1 := curr.Next
		node2 := curr.Next.Next

		// 模拟相邻节点交换
		// temp -> node1 -> node2 交换 temp -> node2 -> node1
		curr.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		// 更新交换节点前一个
		curr = node1

	}
	return dummyHead.Next
}
