/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	var max int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			max++
		} else if max != 0 {
			break
		}
	}
	return max
}

// @lc code=end
