package medium

import "testing"

// 86
func TestPartition(t *testing.T) {
	a := GenerateListByTail([]int{1, 3, 6, 5})
	c := partition(a, 4)
	c.Show()
}
