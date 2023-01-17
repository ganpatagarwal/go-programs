package main

import (
	"log"
)

func main() {
	input := map[int][]int{
		6: {5, 6, 7, 0, 1, 2, 3}, // expected 4
		3: {4, 5, 6, 7, 0, 1, 2}, // expected -1
		2: {1},                   // expected -1
	}

	for target, nums := range input {
		log.Println(search(nums, target))
	}
}

// time - O(N)
// space - O(1)
// linear search
func search1(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

// time - O(log N)
// space - O(1)
// binary search
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		log.Println("mid ", nums[mid])

		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if target >= nums[left] {
			right = mid - 1
		} else {
			left = left + 1
		}
	}

	return -1
}
