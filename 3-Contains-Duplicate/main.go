package main

import (
	"log"
)

func main() {
	nums := []int{1, 2, 3, 4}
	log.Println(containsDuplicate(nums))
}

// time - O(N)
// space - O(N)
func containsDuplicate(nums []int) bool {
	data := map[int]bool{}
	for _, num := range nums {
		if _, exist := data[num]; exist {
			return true
		}

		data[num] = true
	}
	return false
}
