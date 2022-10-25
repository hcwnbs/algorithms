package medium

import "strings"

// 简化路径
func simplifyPath(path string) string {
	var stack []string
	for _, p := range strings.Split(path, "/"){
		if p == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if p != "." && p != "" {
			stack = append(stack, p)
		}
	}

	return "/" + strings.Join(stack, "/")
}
