/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			l++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] == val {
					continue
				}
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return len(nums) - l
}

// @lc code=end
