package main

import (
	"log"
)

func main() {
	inputs := map[int]string{
		2: "ABAB", //4
		//2: "BAAAB",   //5
		1: "AABABBA", //4
		3: "ABCDE",   //4
		0: "ABAA",    // 2
	}

	for k, s := range inputs {
		log.Println(s, characterReplacement(s, k))
	}
}

// time - O(N)
// space - O(1) - since we will store at max 26 alphabets
// brute force - incomplete
func characterReplacement(s string, k int) int {
	charMap := map[rune]int{}
	res := 0
	rs := []rune(s)

	// initialize all letters with 0 count
	for i := 0; i < len(rs); i++ {
		charMap[rs[i]] = 0
	}

	l := 0
	for r := 0; r < len(rs); r++ {
		charMap[rs[r]] = charMap[rs[r]] + 1
		for (r-l+1)-maxCharCount(charMap) > k { // if number of required replacements is grater than allowed
			charMap[rs[l]] = charMap[rs[l]] - 1
			l += 1
		}

		res = max(res, r-l+1)
	}

	return res
}

func maxCharCount(charMap map[rune]int) int {
	m := 0
	for _, v := range charMap {
		if v > m {
			m = v
		}
	}
	return m
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
