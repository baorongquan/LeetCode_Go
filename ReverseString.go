// LeetCode 344. Reverse String

/*Write a function that takes a string as input and returns the string reversed.

Example:
Given s = "hello", return "olleh".*/

package main

import "fmt"

func main() {
	fmt.Println(reverseString("Hello, 世界"))
}

func reverseString(s string) string {
	rs := []rune(s) // rs := []byte(s) 不能用于Unicode字符
	for i, j := 0, len(rs)-1; i < len(rs)/2; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
