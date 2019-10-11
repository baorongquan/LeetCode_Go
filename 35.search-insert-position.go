/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 0
	low := 0
	high := len(nums)
	for {
		index = (high + low) / 2
		if nums[index] == target {
			break
		} else if nums[index] < target {
			if index+1 == len(nums) || nums[index+1] > target {
				index++
				break
			}
			low = index + 1
		} else {
			if index-1 < 0 || nums[index-1] < target {
				break
			}
			high = index - 1
		}
	}
	return index
}

// @lc code=end
