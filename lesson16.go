// usr/bin/go
// author: Ali
// date  : 12/10/2023
package main

import (
	"fmt"
)

type Gopher struct {
	name string
	age  int
}

func main() {
	// Go dasturlash tilida structlar
	gopher := Gopher{name: "Hello", age: 18}
	fmt.Println("Name:", gopher.name)
	fmt.Println("Age:", gopher.age)
}
