package main

import (
	"log"
)

func main() {
	input := map[int][]int{
		3: {5, 6, 7, 0, 1, 2, 3}, // expected 6
		0: {4, 5, 6, 7, 0, 1, 2}, // expected 4
		2: {1},                   // expected -1
		5: {5, 1, 2, 3, 4},       // expected 0
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
		//log.Println("mid ", nums[mid])

		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] { // if left subarray is sorted
			if target >= nums[left] && target < nums[mid] { // if target sits in left sorted subarray
				right = mid - 1
			} else { // target is in right subarray
				left = mid + 1
			}
			// if left subarray is not sorted, then assume that right subarray is sorted
		} else if target < nums[mid] || target > nums[right] { // target is smaller than mid-value or target is greater than last value of right subarray
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
