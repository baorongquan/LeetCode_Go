/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	l := 0
	for i := 0; i < len(nums)-l; i++ {
		if nums[i] == val {
			l++
			for j := len(nums) - l; j > i; j-- {
				if nums[j] == val {
					l++
					continue
				} else {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
	return len(nums) - l
}

// @lc code=end
