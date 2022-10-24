package easy

// 分割数组
func partitionDisjoint(nums []int) int {
	maxLeft := nums[0]
	maxAll := nums[0]
	result := 0
	n := len(nums)
	for i := 1; i < n-1; i++ {
		if nums[i] > maxAll {
			maxAll = nums[i]
		}

		if nums[i] >= maxLeft {
			continue
		}

		maxLeft = maxAll
		result = i
	}
	return result + 1
}

// 第一次做法
//func partitionDisjoint(nums []int) int {
//	l := len(nums)
//	var max int
//	for i := 0; i < l; {
//		for j := i+1; j < l; j++ {
//			if nums[j] >= nums[i] {
//				max = j
//				break
//			}
//		}
//		for k := l-1; k >= max; k-- {
//			if nums[k] <= nums[max] {
//				if k == max {
//					return max
//				}
//				if nums[k] >= nums[i] {
//					continue
//				} else {
//					i = max
//					break
//				}
//			}
//		}
//	}
//	return max
//}
