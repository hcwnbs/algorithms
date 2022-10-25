package medium

// 反转链表 II
// [1, 2, 3, 4, 5] 2, 4
// -> [1, 3, 2, 4, 5] -> [1, 4, 3, 2, 5]
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var prev = head
	curr := head
	i := 1
	for curr != nil {
		if i < left-1 {
			curr = curr.Next
			i++
			continue
		}

		if i == left-1 {
			prev = curr
			curr = curr.Next
			i++
			continue
		}

		if i == right {
			return head
		}

		next := curr.Next
		curr.Next = curr.Next.Next
		if prev == head && left == 1 {
			next.Next = prev
			head = next
			prev = head
		} else {
			next.Next = prev.Next
			prev.Next = next
		}
		i++
	}
	return head
}
