package main

import (
	"fmt"
)

var cached = make(map[uint64]uint64)

func fibonacci(n uint64) uint64 {
	if n < 2 {
		return n
	}
	if val, exist := cached[n]; exist {
		return val
	}

	val := fibonacci(n-1) + fibonacci(n-2)
	cached[n] = val
	return val
}

func main() {
	fmt.Println(fibonacci(98))
	fmt.Println(fibonacci(99))
	fmt.Println(fibonacci(100))
	fmt.Println(fibonacci(101))

	//fmt.Println(fibonacci(0))
	//fmt.Println(fibonacci(1))
	//fmt.Println(fibonacci(2))
	//fmt.Println(fibonacci(3))
	//fmt.Println(fibonacci(4))
	//fmt.Println(fibonacci(5))
	//fmt.Println(fibonacci(6))
	//fmt.Println(fibonacci(7))
	//fmt.Println(fibonacci(8))
	//fmt.Println(fibonacci(9))

}
