package main

import (
	"log"
)

func main() {
	input := map[int][]int{
		1: {1},                         // expected [1]
		2: {1, 1, 1, 2, 2, 3},          // expected [1,2]
		3: {4, 1, 1, 4, 2, 2, 2, 3, 4}, // expected [1,2,4]
	}

	for topK, nums := range input {
		log.Println(topKFrequent(nums, topK))
	}
}

// time - 0(n)
// space - 0(n)
func topKFrequent(nums []int, k int) []int {
	countMap := map[int]int{}

	// count each num occurrence and save it to map
	for _, n := range nums {
		val, exist := countMap[n]
		if exist {
			countMap[n] = val + 1
		} else {
			countMap[n] = 1
		}
	}

	// create a slice which will store list of values with a particular count
	// if all the nums in array are same, this slice will have last entry with len(nums)
	type data []int
	d := make([]data, len(nums)+1)

	// populate slice with data from count map
	for ki, vi := range countMap {
		d[vi] = append(d[vi], ki)
	}

	var ret []int

	// iterate the slice from reverse and find top k elements
	for i := len(d) - 1; i > 0; i-- {
		if len(ret) == k {
			break
		}

		ret = append(ret, d[i]...)
	}

	return ret

}
