package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	//prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}

// time - O(N)
// space - O(1)
func maxProfit(prices []int) int {
	min := prices[0]
	maxEarning := 0

	for i := 1; i < len(prices); i++ {
		val := prices[i]
		if val > prices[i-1] {
			earning := val - min
			if earning > maxEarning {
				maxEarning = earning
			}
		}

		if val < min {
			min = val
		}

		fmt.Println(min, maxEarning)
	}

	return maxEarning
}
