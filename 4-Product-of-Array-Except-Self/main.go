package main

import (
	"log"
)

func main() {
	//nums := []int{1, 2, 3, 4}
	nums := []int{-1, 1, 0, -3, 3}
	log.Println(productExceptSelf3(nums))
}

// time - O(N*2)
// space - O(1)
func productExceptSelf1(nums []int) []int {
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		p := 1
		for j := 0; j < len(nums); j++ {
			if i != j {
				p = p * nums[j]
			}
		}
		result[i] = p
	}

	return result
}

// time - O(N) - with division
// space - O(1)
func productExceptSelf2(nums []int) []int {
	result := make([]int, len(nums))

	totalProduct := 1
	for i := 0; i < len(nums); i++ {
		totalProduct = totalProduct * nums[i]
	}

	for i := 0; i < len(nums); i++ {
		result[i] = totalProduct / nums[i]
	}

	return result
}

// time - O(N)
// space - O(1)
func productExceptSelf3(nums []int) []int {
	result := make([]int, len(nums))

	// prefix products
	pre := 1
	for i := 0; i < len(nums); i++ {
		result[i] = pre
		pre = pre * nums[i]
	}

	// postfix products
	post := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * post
		post = post * nums[i]
	}

	return result
}
