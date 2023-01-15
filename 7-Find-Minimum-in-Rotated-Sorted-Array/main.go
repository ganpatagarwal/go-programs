package main

import (
	"log"
)

func main() {
	input := [][]int{
		{3, 4, 5, 1, 2},       // expected 1
		{4, 5, 6, 7, 0, 1, 2}, // expected 0
		{11, 13, 15, 17},      // expected 11
	}

	for _, nums := range input {
		log.Println(findMin1(nums))
	}
}

// time - O(N)
// space - O(1)
// linear
func findMin(nums []int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	return min
}

// time - O(log N)
// space - O(1)
// binary search
func findMin1(nums []int) int {
	left := 0
	right := len(nums) - 1
	m := nums[0]

	for left <= right {
		if nums[left] < nums[right] {
			return min(m, nums[left])
		}

		mid := (left + right) / 2
		log.Println("mid ", nums[mid])
		m = min(m, nums[mid])

		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return m
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
