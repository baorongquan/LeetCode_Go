import "strconv"

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	result := "1"
	if n > 1 {
		for i := 1; i < n; i++ {
			preSay := make([]byte, 0)
			count := 1
			b := result[0]
			for i := 1; i < len(result); i++ {
				if result[i] == b {
					count++
				} else {
					preSay = append(preSay, []byte(strconv.Itoa(count))...)
					preSay = append(preSay, b)
					b = result[i]
					count = 1
				}
			}
			preSay = append(preSay, []byte(strconv.Itoa(count))...)
			preSay = append(preSay, b)
			result = string(preSay)
		}
	}
	return result
}

// @lc code=end
