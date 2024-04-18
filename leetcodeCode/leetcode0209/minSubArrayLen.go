package leetcode0209

import "math"

// minSubArrayLen
// leetcode 209 取长度最小的子数组
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// 快慢双指针
	l, r := 0, 0
	// 当前最小长度
	length := math.MaxInt
	// 总数
	sum := 0

	for r < n {
		sum += nums[r]
		for sum >= target {
			length = min(length, r-l+1)
			// 减去慢指针的对应的值，并慢指针后移
			sum -= nums[l]
			l++
		}
		// 快指针后移
		r++
	}
	if length == math.MaxInt {
		return 0
	}
	return length
}

// min 取两者最小值
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
