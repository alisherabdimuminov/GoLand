// usr/bin/go
// author: Ali
// date  : 27/09/2023
package main

import (
	"fmt"
)

func main() {
	// Go dasturlash tilida map
	// var var_name = make(map[key_type]val_type)
	var langs = make(map[int]string)
	langs[1] = "Python"
	langs[2] = "C"
	langs[3] = "C++"
	langs[4] = "Go"
	langs[5] = "Rust"
	langs[6] = "Zig"
	fmt.Println(langs)

	// key orqali value ni olish
	fmt.Println(langs[4])

	// langs[key] -> bu bizga ikkita qiymat qaytaradi. 1 - value (agar key mavjud bulsa), 2 - true (agar key mavjud bulsa)/false
	python, isExists := langs[1]
	fmt.Println("Language:", python, "Mavjudmi:", isExists)

	// map dan o'chirish
	delete(langs, 1) // langs dan key=1 bulgan element o'chib ketadi.
	fmt.Println(langs)

	// map ni tozalash
	clear(langs)
	fmt.Println(langs)
}
