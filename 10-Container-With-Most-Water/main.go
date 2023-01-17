package main

import (
	"log"
)

func main() {
	inputs := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7}, //49
		{1, 1},                      //1
		{1, 2, 1},                   // 2
	}

	for _, nums := range inputs {
		log.Println(maxArea1(nums))
	}
}

// time - O(N*2)
// space - O(1)
// brute force
func maxArea1(height []int) int {
	maxA := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			width := j - i
			a := min(height[i], height[j]) * width
			maxA = max(a, maxA)
		}
	}
	return maxA
}

// time - O(N)
// space - O(1)
// two pointer method
func maxArea(height []int) int {
	maxA := 0
	left := 0
	right := len(height) - 1

	for left < right {
		width := right - left
		h := min(height[left], height[right])
		maxA = max(maxA, width*h)
		if height[left] <= height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return maxA
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
