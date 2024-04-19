package leetcode0707

// ListNode 指针
type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	head *ListNode
	// 记录链表长度
	Size int
}

func Constructor() MyLinkedList {
	//哨兵节点作为头节点(虚拟头节点)
	return MyLinkedList{&ListNode{}, 0}
}

func (this *MyLinkedList) Get(index int) int {
	// 无效索引
	if index < 0 || index >= this.Size {
		return -1
	}
	// current 等于真正头节点
	cur := this.head.Next
	for i := 0; i < index; i++ {
		cur = cur.Next // 遍历到索引所在的节点
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// 题目要求
	// index 等于链表的长度，那么该节点会被追加到链表的末尾
	// index 大于链表的长度，该节点不会插入链表中
	// 所以 index 小于 0，这里应该默认index =0
	if index < 0 {
		index = 0
	} else if index > this.Size {
		return
	}

	newNode := &ListNode{Val: val} // 创建新节点
	cur := this.head               // 创建cur 遍历链表
	for i := 0; i < index; i++ {   // 遍历到指定索引的前一个节点
		cur = cur.Next
	}
	newNode.Next = cur.Next // 新节点指向原索引节点
	cur.Next = newNode      // 原索引的前一个节点指向新节点
	this.Size++             // 链表大小增加1
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}

	cur := this.head             // 创建cur 遍历链表
	for i := 0; i < index; i++ { // 遍历到指定索引的前一个节点
		cur = cur.Next
	}
	cur.Next = cur.Next.Next

	this.Size-- // 链表大小减一

}
