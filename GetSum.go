// LeetCode 371. Sum of Two Integers
/*Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.

Example:
Given a = 1 and b = 2, return 3.*/

package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Println(getSum(a, b))
}

func getSum(a, b int) int {
	var num, carry int
	num = a
	carry = b
	for carry != 0 {
		num = a ^ b
		carry = (a & b) << 1
		a = num
		b = carry
	}
	return num
}
