/* Given an array of integers, return indices of the two numbers such that they add up to a specific target.

 * You may assume that each input would have exactly one solution.

 * Example:
 * Given nums = [2, 7, 11, 15], target = 9,

 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 **/

package main

import "fmt"

func main() {
	fmt.Println("two sum")
	var array []int = []int{2, 7, 11, 15}
	var target = 9
	fmt.Println(twoSum(array, target))
}

func twoSum(nums []int, target int) []int {
	len := len(nums)
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}
