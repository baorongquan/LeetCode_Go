// Leetcode 463. Island Perimeter
/*You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water. Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells). The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.

Example:

[[0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]]

Answer: 16*/

package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	fmt.Println(grid)
	fmt.Println(islandPerimeter(grid))
}

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	lenght := len(grid[0])
	hight := len(grid)
	for i := 0; i < hight; i++ {
		for j := 0; j < lenght; j++ {
			if grid[i][j] == 1 {
				edge := 0
				if j == 0 || 0 == grid[i][j-1] { // left
					edge++
				}
				if i == 0 || 0 == grid[i-1][j] { // up
					edge++
				}
				if j == lenght-1 || 0 == grid[i][j+1] { // right
					edge++
				}
				if i == hight-1 || 0 == grid[i+1][j] { // down
					edge++
				}
				perimeter += edge
			}
		}
	}
	return perimeter
}
