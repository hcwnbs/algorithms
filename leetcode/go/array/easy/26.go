package easy

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	r := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			r++
			nums[r] = nums[i]
		}
	}
	return r+1
}
