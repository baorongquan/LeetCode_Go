// LeetCode 104. Maximum Depth of Binary Tree
/*Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// input Tree
//			 1
//			/ \
//		2     3
//	 / \   / \
//	4   5 6   7
//         \
//				  8
// output 4
func main() {
	root := new(TreeNode)
	root.Val = 1
	root.Left = &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}
	root.Right = &TreeNode{3, &TreeNode{6, nil, &TreeNode{8, nil, nil}}, &TreeNode{7, nil, nil}}
	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}
