package easy

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateListByHead(val []int) *ListNode {
	var head *ListNode
	for i := 0; i < len(val); i++ {
		newNode := &ListNode{
			Val:  val[i],
			Next: head,
		}
		head = newNode
	}
	return head
}

func GenerateListByTail(val []int) *ListNode {
	var head = &ListNode{}
	var curr = head
	for i := 0; i < len(val); i++ {
		newNode := &ListNode{
			Val:  val[i],
			Next: nil,
		}
		curr.Next = newNode
		curr = newNode
	}
	return head.Next
}

func (l *ListNode) Show() {
	curr := l
	var val []int
	for curr != nil{
		val = append(val, curr.Val)
		curr = curr.Next
	}
	fmt.Println(val)
}