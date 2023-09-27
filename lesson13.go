// usr/bin/go
// author: Ali
// date  : 27/09/2023
package main

import (
	"fmt"
)

// rekursiya
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// fibonacci
func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// Go dasturlash tilida rekursiya
	// rekursiya nimaligini tushinish uchun rekursiyani tushinish kerak 😂

	// factorial
	var n int = 5
	fac := factorial(n)
	fmt.Println(fac)

	// fibonacci
	var m int = 10
	fib := fibonacci(m)
	fmt.Println(fib)
}
