import "strconv"

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	s := "1"
	if n > 1 {
		for i := 1; i < n; i++ {
			s = countPrev(s)
		}
	}
	return s
}

// countPrev 数之前的字符串
func countPrev(s string) string {
	result := ""
	count := 1
	b := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] == b {
			count++
		} else {
			result += strconv.Itoa(count) + string(b)
			b = s[i]
			count = 1
		}
	}
	result += strconv.Itoa(count) + string(b)
	return result
}

// @lc code=end
