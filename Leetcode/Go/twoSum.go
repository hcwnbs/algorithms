package twoSum

/*
时间复杂度：O(N)
空间复杂度：O(N)

*/
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums{
		if j, ok := hashTable[target - x]; ok{
			return []int{j, i}
		}
		hashTable[x] = i
	}
	return nil
}
