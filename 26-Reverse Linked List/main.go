package main

import (
	"log"
)

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	newHead := reverseList(head)
	for newHead != nil {
		log.Println(newHead.Val)
		newHead = newHead.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// iterative
// time - O(n)
// space - O(1)
func reverseList(head *ListNode) *ListNode {
	var prev, curr *ListNode

	prev = nil
	curr = head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
