package leetcode0242

// isAnagram
// leetcode 242. 有效的字母异位词
func isAnagram(s string, t string) bool {
	// 如果 s,t 互为字母异位词 ,那么两个长度一定相等
	if len(s) != len(t) {
		return false
	}

	cnt := make(map[rune]int)
	for _, ch := range s {
		cnt[ch]++
	}

	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}
