package leetcode0019

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: head}
	// 快指针比慢指针间隔n,但快指针指向nil时，慢指针指向倒数第n个元素的前一个
	low, fast := dummyHead, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 快指针,同时走
	for ; fast != nil; fast = fast.Next {
		low = low.Next
	}
	low.Next = low.Next.Next
	return dummyHead.Next
}
