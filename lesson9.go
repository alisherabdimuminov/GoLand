// usr/bin/go
// author: Ali
// date  : 25/09/2023
package main

import (
	"fmt"
	"slices"
)

func main() {
	// remove element from array/slice
	// method 1: slices yordamida
	var numbers = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(numbers)
	numbers = slices.Delete(numbers, 1, 2) // arguments: slice, i, j -> i dan j gacha bo'lgan elementlarni o'chiradi.
	fmt.Println(numbers)
	// n - elementni o'chirish. misol uchun 4
	var primes = []int{2, 3, 5, 7, 9, 11}
	fmt.Println("Tub sonlar:", primes)
	primes = slices.Delete(primes, 4, 5)
	fmt.Println("Tub sonlar:", primes)

	// method 2: append yordamida
	// arrayga arrayning i-elementigacha va undan keyingilarni qo'shadi. i-elementni tashlab ketadi.
	var evens = []int{2, 4, 6, 7, 8, 10}
	fmt.Println("Xato juft sonlar", evens)
	var remIndex int = 3
	element := evens[remIndex]

	evens = append(evens[:remIndex], evens[remIndex+1:]...)

	fmt.Println("O'chirilishi kerak bo'lgan son:", element)
	fmt.Println("Juft sonlar:", evens)
}
