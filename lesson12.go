// usr/bin/go
// author: Ali
// date  : 27/09/2023
package main

import (
	"fmt"
)

// argumentsiz va returnsin
func salom(name string) {
	fmt.Println("Salom " + name)
}

// argument/return bilan
func change(a *int) {
	*a += 1
}

func add(a, b int) int {
	return a + b
}

// istalgancha argument va return

func sum(numbers ...int) int {
	answer := 0
	for _, number := range numbers {
		answer += number
	}
	return answer
}

func main() {
	// Go dasturlash tilida functions
	// sintaksis: func func_name(arguments)return_type {body}
	salom("Ali")

	var a int = 230
	var b int = 120
	var c int = add(a, b)
	fmt.Println(c)

	// istalgancha argument va return
	num1 := 12
	num2 := 34
	num3 := 90
	num4 := 45
	res := sum(num1, num2, num3, num4)
	fmt.Println(res)

	// yuqoridagi sum funksiyamizga array ni ham quyidagicha uzatsak bo'ladi(arraydan keyin 3 ta nuqta bilan)
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arraySum := sum(numbers...)
	fmt.Println(arraySum)
}
