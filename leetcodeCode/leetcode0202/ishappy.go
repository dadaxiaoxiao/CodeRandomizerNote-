package leetcode0202

// isHappy
// leetcode 202. 快乐数
func isHappy(n int) bool {
	set := make(map[int]bool)
	// 循环跳出条件 n = 1 或者 无限循环（这里就是 set[n] == true, 哈希表存在元素）
	for n != 1 && !set[n] {
		n, set[n] = getSum(n), true
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	// 每个数字的平方和
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
