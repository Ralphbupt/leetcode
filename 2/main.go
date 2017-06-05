package main

import "fmt"

func main() {
	l1 := makeList(5432343333332222223)
	l2 := makeList(11111)
	l := addTwoNumbers(l1, l2)
	for l != nil {
		fmt.Print(l.Val)
		l = l.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	pos := 0
	for l1 != nil || l2 != nil || pos!=0{
		vx, vy := 0, 0
		if l1 != nil {
			vx = l1.Val
		}
		if l2 != nil {
			vy = l2.Val
		}
		tmp := vx + vy + pos
		val := tmp % 10
		pos = tmp / 10
		node := &ListNode{val, nil}
		cur.Next = node
		cur = cur.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}
	return head.Next
}

func makeList(num int64) *ListNode {
	h := &ListNode{}
	c := h
	for num > 0 {
		val := num % 10
		num = num / 10
		node := &ListNode{int(val), nil}
		c.Next = node
		c = c.Next
	}
	return h.Next
}
