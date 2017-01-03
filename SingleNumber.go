// LeetCode 136. SingleNumber
/*Given an array of integers, every element appears twice except for one. Find that single one.

Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?*/

package main

import "fmt"

func main() {
	nums := []int{11, 3, 4, 3, 11, 4, 5}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	singleNumber := 0
	for i := 0; i < len(nums); i++ {
		singleNumber ^= nums[i]
	}
	return singleNumber
}
