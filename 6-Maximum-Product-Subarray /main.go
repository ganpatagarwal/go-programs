package main

import (
	"log"
)

func main() {
	input := [][]int{
		{2, 3, -2, 4},      // expected 6
		{0, 2},             // expected 2
		{-3, -1, -1},       // expected 3
		{3, -1, 4},         // expected 4
		{-2, 3, -4},        // expected 24
		{2, -5, -2, -4, 3}, // expected 24
		{-2, 1, 0},         // expected 1
	}

	for _, nums := range input {
		log.Println(maxProduct1(nums))
	}
}

// time - O(N)
// space - O(1)
// dynamic programming
func maxProduct(nums []int) int {
	maxP := nums[0]
	curMax, curMin := 1, 1

	for _, n := range nums {
		if n == 0 {
			curMax, curMin = 1, 1
			continue
		}
		tmpMax := curMax * n
		curMax = max(max(curMin*n, curMax*n), n)
		curMin = min(max(curMin*n, tmpMax), n)
		maxP = max(maxP, curMax)
	}

	return maxP
}

// time - O(N*2)
// space - O(1)
// brute force
func maxProduct1(nums []int) int {
	maxP := nums[0]
	lenNums := len(nums)
	for i := 0; i < lenNums; i++ {
		currP := nums[i]
		for j := i + 1; j < lenNums; j++ {
			maxP = max(maxP, currP)
			currP *= nums[j]
		}
		maxP = max(maxP, currP)
	}

	return maxP
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
