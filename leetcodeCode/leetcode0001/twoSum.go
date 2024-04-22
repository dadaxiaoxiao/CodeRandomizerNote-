package leetcode0001

// twoSum
// leetcode 1. 两数之和
func twoSum(nums []int, target int) []int {
	hashtable := make(map[int]int)
	for index, v := range nums {
		if preIndex, ok := hashtable[target-v]; ok {
			return []int{preIndex, index}
		}
		hashtable[v] = index
	}
	return nil
}
