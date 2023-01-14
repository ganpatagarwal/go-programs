package main

import (
	"log"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	log.Println(maxSubArray(nums))
}

// time - O(N)
// space - O(1)
func maxSubArray(nums []int) int {

	maxSum := nums[0]
	currSum := 0

	for i := 0; i < len(nums); i++ {
		if currSum < 0 {
			currSum = 0
		}

		currSum += nums[i]
		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}
