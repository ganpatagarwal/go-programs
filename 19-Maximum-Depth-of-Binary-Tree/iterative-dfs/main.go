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

type data struct {
	node  *TreeNode
	depth int
}

func maxDepth(root *TreeNode) int {
	result := 0
	stack := []data{
		{
			node:  root,
			depth: 1,
		},
	}

	for len(stack) > 0 {
		lastIndex := len(stack) - 1
		d := stack[lastIndex]
		stack = stack[:lastIndex]

		if d.node != nil {
			log.Println(d.node.Val)
			result = max(result, d.depth)
			stack = append(stack, data{
				node:  d.node.Right,
				depth: d.depth + 1,
			})
			stack = append(stack, data{
				node:  d.node.Left,
				depth: d.depth + 1,
			})
		}
	}
	return result
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
