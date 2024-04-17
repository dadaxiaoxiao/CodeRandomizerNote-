package leetcode0704

// search 二分查找
// 这里使用坐闭右开原则
func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] > target {
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

// search2 二分查找 第二种写法
// 这个是右闭右闭
func search2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] > target {
			// 进入左区间
			right = middle - 1
		} else if nums[middle] < target {
			// 进入右区间
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
