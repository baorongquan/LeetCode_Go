/*Initially, there is a Robot at position (0, 0). Given a sequence of its moves, judge if this robot makes a circle, which means it moves back to the original place.

The move sequence is represented by a string. And each move is represent by a character. The valid robot moves are R (Right), L (Left), U (Up) and D (down). The output should be true or false representing whether the robot makes a circle.

Example 1:
Input: "UD"
Output: true
Example 2:
Input: "LL"
Output: false
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	moves := "UD"
	isCircle := judgeCircle(moves)
	fmt.Println("isCircle :", isCircle)
}

func judgeCircle(moves string) bool {
	return strings.Count(moves, "L") == strings.Count(moves, "R") && strings.Count(moves, "U") == strings.Count(moves, "D")
}
func judgeCircle2(moves string) bool {
	UpDown, LeftRight := 0, 0
	for _, move := range moves {
		switch string(move) {
		case "U":
			UpDown++
		case "D":
			UpDown--
		case "L":
			LeftRight++
		case "R":
			LeftRight--
		}
	}
	if 0 == UpDown && 0 == LeftRight {
		return true
	}
	return false
}
