package leetcode0206

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
// reverseList
// 链表反转
// 这里使用 cur ,pre 双指针法
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	// 新建 cur 暂存节点用来遍历
	cur := head
	for cur != nil {
		// 下一个节点
		next := cur.Next
		// 指针域反转
		cur.Next = pre
		// 先移动 pre,后移动cur
		pre = cur
		cur = next
	}
	return pre
}
*/

// reverseList
// 链表反转
// 这里使用递归
func reverseList(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	// 下一个指针节点
	next := head.Next
	// 指针域反转
	head.Next = pre
	return reverse(head, next)
}
