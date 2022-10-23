package medium

// 环形链表 II
// 快慢指针
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			fast = head
			for {
				if slow == fast {
					return slow
				}
				slow = slow.Next
				fast = fast.Next
			}
		}
	}
	return nil
}

// 哈希表
//func detectCycle(head *ListNode) *ListNode {
//	var list = make(map[*ListNode]bool)
//	for head != nil {
//		if list[head] {
//			return head
//		}
//		list[head] = true
//		head = head.Next
//	}
//	return nil
//}
