package easy

// 交替合并字符串
func mergeAlternately(word1 string, word2 string) string {
	n, m := len(word1), len(word2)
	var result = make([]byte, 0, m+n)
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < n {
			result = append(result, word1[i])
		}

		if i < m {
			result = append(result, word2[i])
		}
	}
	return string(result)
}

//递归做法
//func mergeAlternately(word1 string, word2 string) string {
//	if word1 == "" {
//		return word2
//	}
//	if word2 == "" {
//		return word1
//	}
//	var i = 0
//	newWord := string(word1[i]) + string(word2[i])
//	i++
//	return newWord + mergeAlternately(word1[i:], word2[i:])
//}
