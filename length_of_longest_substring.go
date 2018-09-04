package leetcode

// 3. 无重复字符的最长子串
// 给定一个字符串，找出不含有重复字符的最长子串的长度。

// 难度：中等

// 示例1:
// 输入: "abcabcbb"
// 输出: 3
// 解释: 无重复字符的最长子串是 "abc"，其长度为 3。

// 示例 2:
// 输入: "bbbbb"
// 输出: 1
// 解释: 无重复字符的最长子串是 "b"，其长度为 1。

// 示例 3:
// 输入: "pwwkew"
// 输出: 3
// 解释: 无重复字符的最长子串是 "wke"，其长度为 3。
// 请注意，答案必须是一个子串，"pwke" 是一个子序列 而不是子串。

// 暴力法
// 时间复杂度：O(n^2)
// 执行用时：460 ms
func LengthOfLongestSubstring1(s string) int {
	var c []int
	b := []byte(s)
	m := make(map[byte]int)
	count := 0
	for i := 0; i < len(b); i++ {
		count++
		m[b[i]] = i
		for j := i + 1; j < len(b); j++ {
			if _, ok := m[b[j]]; ok {
				c = append(c, count)
				count = 0
				m = make(map[byte]int)
				break
			}
			count++
			m[b[j]] = j
		}
		c = append(c, count)
		count = 0
		m = make(map[byte]int)
	}
	if len(c) > 0 {
		for _, num := range c {
			if num > count {
				count = num
			}
		}
		return count
	}
	return count
}
