/** You are given two linked lists representing two non-negative numbers.
 *  The digits are stored in reverse order and each of their nodes contain a single digit.
 *  Add the two numbers and return it as a linked list.
 *  Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 *  Output: 7 -> 0 -> 8
 **/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var a1 []int = []int{3, 9, 9, 9}
	var a2 []int = []int{7}
	var l1 *ListNode = new(ListNode)
	var l2 *ListNode = new(ListNode)
	tl1 := l1
	tl2 := l2
	for _, v := range a1 {
		var t = new(ListNode)
		t.Val = v
		l1.Next = t
		l1 = l1.Next
	}
	for _, v := range a2 {
		var t = new(ListNode)
		t.Val = v
		l2.Next = t
		l2 = l2.Next
	}
	fmt.Println("l1", l1)
	for tt := tl1; tt != nil; tt = tt.Next {
		fmt.Print(tt.Val)
	}
	fmt.Println("tl1 ", tl1.Next)
	result := addTwoNumbers(tl1.Next, tl2.Next)
	for result != nil {
		fmt.Print(result.Val, " $")
		result = result.Next
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s int = 0
	var pL1 *ListNode
	var pL2 *ListNode
	tmpL1 := l1
	for l1 != nil && l2 != nil {
		l1.Val += s
		s = (l1.Val + l2.Val) / 10
		l1.Val = (l1.Val + l2.Val) % 10
		pL1 = l1
		pL2 = l2
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil && l2 != nil {
		pL1.Next = pL2.Next
	}

	if s == 1 {
		ppL1 := pL1
		pL1 = pL1.Next
		for pL1 != nil && s == 1 {
			ts := (pL1.Val + s) / 10
			pL1.Val = (pL1.Val + s) % 10
			s = ts
			ppL1 = pL1
			pL1 = pL1.Next
		}
		if s == 1 {
			var node *ListNode = new(ListNode)
			node.Val = s
			ppL1.Next = node
		}
	}

	return tmpL1
}
