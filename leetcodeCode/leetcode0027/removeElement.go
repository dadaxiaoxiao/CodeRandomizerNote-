package leetcode00027

// removeElement 移除元素
// 这里使用双指针，并对 nums 进行了缩容
func removeElement(nums []int, val int) int {
	res := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	nums = nums[:res]
	return res
}
