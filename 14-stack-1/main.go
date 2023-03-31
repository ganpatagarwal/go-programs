package main

import (
	"log"
)

type node struct {
	val  int
	prev *node
}

type stack struct {
	current *node
}

func (s *stack) add(val int) {
	n := &node{
		val:  val,
		prev: nil,
	}

	if s.current != nil {
		n.prev = s.current
	}

	s.current = n
}

func (s *stack) remove() int {
	if s.current == nil {
		return -1
	}

	val := s.current.val
	s.current = s.current.prev

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
