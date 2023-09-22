// usr/bin/go
// author: Ali
// date  : 22/09/2023
package main

import "fmt"

func main() {
	// Go dasturlash tilida IOStream

	var gopher string = "Gopher"
	var pythonist string = "Pythonist"

	// output stream
	fmt.Print("Hello, Gophers!")                        // konsolega chiqarish
	fmt.Println("Hello, Pythonists")                    // yangi qatordan konsolega chiqarish
	fmt.Printf("Hello %s\nBye %s\n", gopher, pythonist) // format yordamida konsolega chiqarish
	// formatlash turlari
	/*
		%d -> Integer
		%c -> Character
		%f -> Float ( %.[num]f num xona aniqlikda formatlash )
		%s -> String
		%o -> Base 8
		%x -> Base 16
	*/
	var name string = "Ali"
	var age int = 18
	var PI float32 = 3.1415
	fmt.Printf("Hello, %s\n", name)
	fmt.Printf("I am %d years old.\n", age)
	fmt.Printf("PI is %f\n", PI)
	fmt.Printf("PI is %.2f\n", PI)

	// input stream
	var username string
	var password int
	fmt.Scan(&username, &password)
	fmt.Printf("Username: %s\nPassword: %d\n", username, password)

	var color string
	fmt.Scanf("%s", &color) // formatli kiritish oqimi
	fmt.Println("Sizning yoqtirgan rangingiz:", color)
}
