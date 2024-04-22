package leetcod0349

// intersection
// leetcode  349. 两个数组的交集
func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	set := make(map[int]struct{}, 0)

	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
			// 哈希表删除，下一次就会判断不存在，这里保证了唯一性
			delete(set, v)
		}
	}
	return res
}
