/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	index := -1
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] && i+len(needle) <= len(haystack) {
			index = i
			for j := 0; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					index = -1
					break
				}
			}
		}
		if index != -1 {
			break
		}
	}
	return index
}

// @lc code=end
