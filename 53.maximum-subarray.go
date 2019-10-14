/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var premax, max int
	premax = nums[0]
	for i := range nums {
		if max+nums[i] > 0 {
			if max > 0 {
				max += nums[i]
			} else {
				max = nums[i]
			}
		} else {
			if premax <= 0 {
				max = nums[i]
			} else {
				max = 0
			}
		}
		if max > premax {
			premax = max
		}
	}
	return premax
}

// @lc code=end
