/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	var flag bool
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 > 9 {
			digits[i] = 0
			if i == 0 {
				flag = true
			}
		} else {
			digits[i] += 1
			break
		}
	}

	if !flag {
		return digits
	} else {
		result := make([]int, 0, 1+len(digits))
		result = append(result, 1)
		result = append(result, digits...)
		return result
	}
}

// @lc code=end
