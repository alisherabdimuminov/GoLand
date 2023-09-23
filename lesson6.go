// usr/bin/go
// author: Ali
// date  : 23/09/2023
package main

import "fmt"

func main() {
	// Go dasturlash tilida shart operatorlari
	// if/else
	var number int
	fmt.Print("N sonini kiriting: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println(number, "juft son.")
	} else {
		fmt.Println(number, "toq son.")
	}

	// if/else/else if
	var language string
	fmt.Print("Dasturlash tilini tanlang[c, c++, go, python]")
	fmt.Scan(&language)
	if language == "c" {
		fmt.Println("Vapshe very language. Father of all languages")
	} else if language == "c++" {
		fmt.Println("Brother of Gophers")
	} else if language == "go" {
		fmt.Println("Meni aytayabdi")
	} else if language == "python" {
		fmt.Println("nice))")
	} else {
		fmt.Println("Bu tilni bilmayman.")
	}
}
