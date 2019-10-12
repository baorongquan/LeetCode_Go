import (
	"bytes"
	"strconv"
)

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	result := "1"
	preSay := bytes.NewBuffer(make([]byte, 0, 5))
	if n > 1 {
		for i := 1; i < n; i++ {
			preSay.Reset()
			count := 1
			b := result[0]
			for i := 1; i < len(result); i++ {
				if result[i] == b {
					count++
				} else {
					preSay.WriteString(strconv.Itoa(count))
					preSay.WriteByte(b)
					b = result[i]
					count = 1
				}
			}
			preSay.WriteString(strconv.Itoa(count))
			preSay.WriteByte(b)
			result = preSay.String()
		}
	}
	return result
}

// @lc code=end
