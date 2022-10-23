package easy

import (
	"fmt"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	a := "test"
	b := "ddd"
	c := mergeAlternately(a, b)
	fmt.Println(c)
}
