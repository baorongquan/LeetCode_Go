/* Given a string, find the length of the longest substring without repeating characters.

 * Examples:

 * Given "abcabcbb", the answer is "abc", which the length is 3.

 * Given "bbbbb", the answer is "b", with the length of 1.

 * Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

 * Subscribe to see which companies asked this question
 **/
package main

import "fmt"

func main() {
	len := lengthOfLongestSubstring("pwwkew")
	fmt.Println(" len ", len)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var max int = 0
	tempMap := make(map[string]int)
	for i, j := 0, 0; i < len(s); i++ {
		if val, ok := tempMap[string(s[i])]; ok {
			if j < val+1 {
				j = val + 1
			}
		}
		tempMap[string(s[i])] = i
		if max < i-j+1 {
			max = i - j + 1
		}
	}
	return max
}
