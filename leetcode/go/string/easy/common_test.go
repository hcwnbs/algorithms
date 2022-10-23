package easy

import (
	"fmt"
	"testing"
)

// 1786
func TestMergeAlternately(t *testing.T) {
	a := "test"
	b := "ddd"
	c := mergeAlternately(a, b)
	fmt.Println(c)
}
