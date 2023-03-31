package main

import (
	"log"
)

type queue struct {
	nums []int
}

func (q *queue) add(val int) {
	q.nums = append(q.nums, val)
}

func (q *queue) remove() int {
	qLen := len(q.nums)

	if qLen == 0 {
		return -1
	}

	if qLen == 1 {
		val := q.nums[0]
		q.nums = []int{}
		return val
	}

	val := q.nums[0]
	q.nums = q.nums[1:]

	return val

}
func main() {
	q := new(queue)

	q.add(3)
	q.add(1)
	q.add(2)

	log.Println(q.remove())
	log.Println(q.remove())
	log.Println(q.remove())
	log.Println(q.remove())
}
