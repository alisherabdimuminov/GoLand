// usr/bin/go
// author: Ali
// date  : 25/09/2023
package main

import "fmt"

func main() {
	// Go dasturlash tilida arraylar bilan ishlash
	// arrayni e'lon qilish
	// 1) var var_name [values_count]values_type
	// 2) var var_name = [values_count]values_type{values}
	// 3) var var_name = make([]int, 10)

	// 1)
	var array1 []int
	fmt.Println(array1)

	// 2)
	var array2 = []int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	// 3)
	var array3 = make([]int, 10)
	fmt.Println(array3)

	// arrayga element element olish, o'zgartirish, qo'shish.
	// element olish: array[index]
	// ! indekslash 0 dan boshlanadi
	var primes = []int{2, 3, 5, 7, 11}
	fmt.Println("Tub sonlar ro'yxati:", primes)
	fmt.Println("Uchinchi tub son:", primes[2])

	for index, value := range primes {
		fmt.Printf("primes[%d] = %d\n", index, value)
	}

	// elementni o'zgartirish
	primes[0] = 0
	fmt.Println("O'zgartirilgan primes[] =", primes)
	primes[0] = 2

	// element qo'shish - append() funksiyasi
	// array = append(array, elements)
	primes = append(primes, 13, 17, 19)
	fmt.Println(primes)

	evens := make([]int, 0, 10)
	i := 0
	for {
		if i > 100 {
			break
		} else {
			if i%2 == 0 {
				evens = append(evens, i)
			}
		}
		i++
	}
	fmt.Println("1 dan 100 gacha juft sonlar:", evens)

	// slices - bo'laklar/qisimlar
	var letters = []string{"A", "l", "i", "s", "h", "e", "r"}
	fmt.Println(letters)

	var ali = letters[0:3] // letters ning 0 - indeksidan 3 - indeksigacha (3 - indeksdagi element olinmaydi)
	fmt.Println(ali)

	var sher = letters[3:]
	fmt.Println(sher)

	// cap - array ning capacity(sig'im) si
	fmt.Printf("cap(letters) = %d\n", cap(letters)) // letters[] ning sig'imi 7 ga teng
	// letters[8] = "b" // runtime error: index out of range [8] with length 7

	// 2d array
	var matrix [3][3]int

	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3

	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	matrix[2][0] = 7
	matrix[2][1] = 8
	matrix[2][2] = 9

	fmt.Println("matrix =", matrix)
}
