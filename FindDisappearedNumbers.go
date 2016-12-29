// LeetCode 448. Find All Numbers Disappeared in an Array

/*Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.*/

/*Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
*/

/* 解题思路：遍历数据，如果数组下标i对应的数num[i]重复的话则通过num[num[i]-1]来标记的时候就会标记到同一个值
   为了在遍历过程中标记并且标记后还可以通过某种方法来获取标记前的值，有两种方法可以参考：
   1. 取模
   2. 取反
	 最后消失了的数(对应下标为disappearednum - 1)我们就标记不到（取模时仍<= len； 取反时仍为正)
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	arr1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers1(arr1))
	arr2 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers2(arr2))
}

func findDisappearedNumbers1(nums []int) []int {
	disNum := make([]int, 0, len(nums))
	len := len(nums)
	for i := 0; i < len; i++ {
		nums[(nums[i]-1)%len] += len
	}
	//fmt.Println(nums)
	for i := 0; i < len; i++ {
		if nums[i] <= len {
			disNum = append(disNum, i+1)
		}
	}
	return disNum
}

func findDisappearedNumbers2(nums []int) []int {
	disNum := make([]int, 0, len(nums))
	len := len(nums)
	for i := 0; i < len; i++ {
		val := int(math.Abs(float64(nums[i])) - 1)
		if nums[val] > 0 {
			nums[val] = -nums[val]
		}
	}
	for i := 0; i < len; i++ {
		if nums[i] > 0 {
			disNum = append(disNum, i+1)
		}
	}
	return disNum
}
