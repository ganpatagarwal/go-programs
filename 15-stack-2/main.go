package main

import (
	"log"
)

type stack struct {
	nums []int
}

func (s *stack) add(val int) {
	s.nums = append(s.nums, val)
}

func (s *stack) remove() int {
	sLen := len(s.nums)
	if sLen == 0 {
		return -1
	}
	val := s.nums[sLen-1]

	s.nums = s.nums[:sLen-1]
	return val
}

func main() {
	s := new(stack)

	s.add(1)
	s.add(3)
	s.add(2)

	log.Println(s.remove())
	log.Println(s.remove())
	log.Println(s.remove())
	log.Println(s.remove())
}
