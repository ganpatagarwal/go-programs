package main

import (
	"log"
)

func main() {
	inputs := map[string]string{
		"ADOBECODEBANC": "ABC", //BANC
		"b":             "b",   //b
		"a":             "aa",  //""
		"aa":            "aa",  //aa
		"bdab":          "ab",  //ab
	}

	for s, t := range inputs {
		log.Println(s, minWindow(s, t))
	}
}

func minWindow(s string, t string) string {
	targetMap := map[rune]int{}
	targetRune := []rune(t)

	for _, target := range targetRune {
		targetMap[target] += 1
	}
	targetLen := len(targetMap)

	sourceRune := []rune(s)
	currentCharMap := map[rune]int{}
	totalLen := 0
	result := []int{-1, -1}
	min := len(s)

	l := 0

	for r := 0; r < len(sourceRune); r++ {
		if val, exist := targetMap[sourceRune[r]]; exist {
			rVal := currentCharMap[sourceRune[r]]
			currentCharMap[sourceRune[r]] = rVal + 1

			if rVal+1 == val {
				totalLen += 1
			}
		}

		for totalLen == targetLen {
			currLen := r - l + 1
			if currLen <= min {
				min = currLen
				result = []int{l, r}
			}

			if val, exist := targetMap[sourceRune[l]]; exist {
				lVal := currentCharMap[sourceRune[l]]
				currentCharMap[sourceRune[l]] = lVal - 1
				if lVal-1 < val {
					totalLen -= 1
				}
			}
			l += 1
		}
	}

	if result[0] == -1 {
		return ""
	}
	return s[result[0] : result[1]+1]

}

// time - O(N*M)
// space - O(1)
// brute force
func minWindow1(s string, t string) string {
	targetMap := map[rune]int{}
	targetRune := []rune(t)

	for _, target := range targetRune {
		targetMap[target] += 1
	}

	sourceRune := []rune(s)
	outputs := []string{}
	currentCharMap := map[rune]int{}

	l := 0

	for r := 0; r < len(sourceRune); r++ {
		if _, exist := targetMap[sourceRune[r]]; !exist && l == r {
			l += 1
		}

		currentCharMap[sourceRune[r]] += 1

		if countFulfilled(targetMap, currentCharMap) {
			outputs = append(outputs, string(sourceRune[l:r+1]))
			log.Println("outputs: ", outputs)
			l = r + 1
			currentCharMap = map[rune]int{}
		}
	}

	min := len(s)
	var result string
	for _, output := range outputs {
		if len(output) <= min {
			min = len(output)
			result = output
		}
	}

	return result
}

func countFulfilled(target, source map[rune]int) bool {
	for k, v := range target {
		if source[k] < v {
			return false
		}
	}
	return true
}
