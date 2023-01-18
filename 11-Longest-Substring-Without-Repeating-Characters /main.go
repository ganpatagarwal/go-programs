package main

import (
	"log"
)

func main() {
	inputs := []string{
		"abcabcabc", //3
		"bbbbb",     //1
		"pwwkew",    //3
		"a",         //1
		"aab",       //2
		"dvdf",      //3
	}

	for _, s := range inputs {
		log.Println(lengthOfLongestSubstring(s))
	}
}

// time - O(N)
// space - O(N)
// sliding window
func lengthOfLongestSubstring(s string) int {
	res := 0
	l := 0

	rs := []rune(s)
	charMap := map[rune]int{}

	for r := 0; r < len(rs); r++ {
		if v, exist := charMap[rs[r]]; exist && v >= l {
			l = v + 1
		}
		res = max(res, r-l+1)
		charMap[rs[r]] = r
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
