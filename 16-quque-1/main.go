package main

import (
	"log"
)

type node struct {
	val  int
	next *node
}

type queue struct {
	first *node
	last  *node
}

func (q *queue) add(val int) {
	n := &node{
		val:  val,
		next: nil,
	}

	if q.last != nil {
		q.last.next = n
		q.last = n
	} else {
		q.first = n
		q.last = n
	}
}

func (q *queue) remove() int {
	if q.first == nil {
		return -1
	}

	val := q.first.val
	q.first = q.first.next

	return val
}

func main() {
	q := new(queue)

	q.add(1)
	q.add(2)
	q.add(3)

	log.Println(q.remove())
	log.Println(q.remove())
	log.Println(q.remove())
	log.Println(q.remove())
}
