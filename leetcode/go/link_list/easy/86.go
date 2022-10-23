package easy

// 分隔链表
func partition(head *ListNode, x int) *ListNode {
	smallHead := &ListNode{}
	LargeHead := &ListNode{}

	currSmall := smallHead
	currLarge := LargeHead

	for head != nil {
		if head.Val < x {
			currSmall.Next = head
			currSmall = head
		} else {
			currLarge.Next = head
			currLarge = head
		}
		head = head.Next
	}

	currLarge.Next = nil
	currSmall.Next = LargeHead.Next
	return smallHead.Next
}
