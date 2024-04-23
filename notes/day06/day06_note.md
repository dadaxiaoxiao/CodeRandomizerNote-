# 代码随想录训练营-Day06

 

## 哈希表相关知识

![](picture/Snipaste_2024-04-23_13-13-48.png)

使用总结

**当我们遇到了要快速判断一个元素是否出现集合里的时候，就要考虑哈希法**。牺牲了空间换取时间。



## LeetCode题目

### 242.有效的字母异词

题目：[242. 有效的字母异位词](https://leetcode.cn/problems/valid-anagram/)

题目截图：

![](picture/Snipaste_2024-04-23_13-16-33.png)



实现代码

```go

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
```



[完整实现代码](https://github.com/dadaxiaoxiao/CodeRandomizerNote-/blob/main/leetcodeCode/leetcode0242/isanagram.go)



### **349. 两个数组的交集** 

题目：[349. 两个数组的交集](https://leetcode.cn/problems/intersection-of-two-arrays/)

题目截图

![](picture/Snipaste_2024-04-23_13-20-46.png)





###  **202. 快乐数** 



![](picture/Snipaste_2024-04-23_13-21-30.png)



###  **1. 两数之和** 