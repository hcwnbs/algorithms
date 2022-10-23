package easy

import (
	"testing"
)

func TestPartition(t *testing.T) {
	a := GenerateListByTail([]int{1, 3, 6, 5})
	c := partition(a, 4)
	c.Show()
}

func TestMergeTwoLists(t *testing.T) {
	a := GenerateListByTail([]int{-10,-10,-9,-4,1,6,6})
	b := GenerateListByTail([]int{-7})
	c := mergeTwoLists(a, b)
	c.Show()
}

//160
func TestGetIntersectionNode(t *testing.T) {
	a := GenerateListByTail([]int{4,1})
	b := GenerateListByTail([]int{5,6,1})
	c := GenerateListByTail([]int{8,4,5})

	curr := a
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = c

	curr = b
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = c

	a.Show()
	b.Show()

	result := getIntersectionNode(a, b)
	result.Show()
}

// 206
func TestReverseList(t *testing.T) {
	var l = &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l.Show()
	n := reverseList(l)
	n.Show()
}

