package main

import (
	"fmt"
	"math/big"
)

func fibonacci(n uint) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}

	a, b := big.NewInt(0), big.NewInt(1)

	for n--; n > 0; n-- {
		a.Add(a, b)
		a, b = b, a
	}

	return b
}

func fibonacci1(n int) int {
	if n < 2 {
		return n
	}

	var a, b int

	b = 1

	// starting from 1 as Fib(1) is already calculated
	for i := 1; i < n; i++ {
		a += b
		a, b = b, a
	}

	return b
}

func main() {
	fmt.Println(fibonacci(100))
}
