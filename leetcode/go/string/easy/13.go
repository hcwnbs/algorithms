package easy

// 罗马数字转整数
func romanToInt(s string) int {
	var result int
	for i := 0; i < len(s); i++ {
		if i + 1 == len(s) {
			result = result + getRomanValue(s[i])
			return result
		}
		if getRomanValue(s[i]) < getRomanValue(s[i+1]) {
			result = result - getRomanValue(s[i])
			continue
		}
		result = result + getRomanValue(s[i])
	}
	return result
}

func getRomanValue(r uint8) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}