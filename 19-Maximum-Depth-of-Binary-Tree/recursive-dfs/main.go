package main

import (
	"log"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	log.Println(root.Val)
	d := max(maxDepth(root.Left), maxDepth(root.Right))
	return d + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	node := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	log.Println(maxDepth(node)) // 3
}
