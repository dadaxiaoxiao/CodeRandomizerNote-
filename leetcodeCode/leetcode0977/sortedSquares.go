package leetcode0977

// sortedSquares
// leetcode997 有序数组的平方
// 这里使用了双指针的思想
func sortedSquares(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	left, right, k := 0, length-1, length-1
	for left <= right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			res[k] = nums[right] * nums[right]
			// right 指针前移动
			right--
		} else {
			res[k] = nums[left] * nums[left]
			// left 指针后移动
			left++
		}
		k--
	}
	return res
}
