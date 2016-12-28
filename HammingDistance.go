// 461.Hamming Distance
//The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

//Given two integers x and y, calculate the Hamming distance.

//Note:
//0 ≤ x, y < 231.
package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(5, 4))
}

func hammingDistance(x, y int) int {
	z := x ^ y
	num := 0
	for z != 0 {
		if z&1 != 0 {
			num++
		}
		z = z >> 1
	}
	return num
}
