// usr/bin/go
// author: Ali
// date  : 29/09/2023
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Go dasturlash tilida random
	// import (... "math/rand")
	var int_random int = rand.Int() // -> int turidagi random son generatsiya qiladi
	fmt.Println(int_random)

	var range_random int = rand.Intn(100) // -> 0 va n oraliqda random son generatsiya qiladi
	fmt.Println(range_random)

	var float_random float32 = rand.Float32() // -> 0.0 va 1.0 oraliqda float32 turidagi random son generatsiya qiladi
	fmt.Println(float_random)

	var permutation []int = rand.Perm(10) // -> n ta elementdan iborat elementlari random sonlar arrayini qaytaradi(0 dan n gacha)
	fmt.Println(permutation)

	var langs = []string{"c", "c++", "go", "python", "rust", "zig", "v"}
	fmt.Println(langs)
	rand.Shuffle(len(langs), func(i, j int) {
		langs[i], langs[j] = langs[j], langs[i]
	}) // -> berilgan arrayning n uzunlikdagi qismini aralashtirib beradi
	fmt.Println(langs)
}
