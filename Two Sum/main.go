package main

import (
	"log"
)

func main() {
	nums := []int{-3, 4, 3, 90}
	target := 0
	log.Println(twoSum(nums, target))
}

// time - O(N)
// space - O(N)
func twoSum(nums []int, target int) []int {
	data := map[int]int{}
	for index, num := range nums {
		if val, exist := data[target-num]; exist {
			return []int{val, index}
		}

		data[num] = index

		log.Println(data)
	}

	return nil
}
