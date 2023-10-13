// usr/bin/go
// author: Ali
// date  : 13/10/2023
package main

import (
	"fmt"
)

// qaytariluvchi qiymat sifatida int qaytaradigan funksiya uzatilyabdi
func generator() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// Go dasturlash tilida generatorcha
	next := generator()

	fmt.Println(next())
	fmt.Println(next())
}
