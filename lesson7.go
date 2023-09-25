// usr/bin/go
// author: Ali
// date  : 25/09/2023
package main

import "fmt"

func main() {
	// Go dasturlash tilida sikl
	// syntax:for i :=0;i < n; i++ {body}
	// namuna 1
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// namuna: while true
	for {
		fmt.Println("go for == c while true")
	}

	// numana: with array
	var langs = []string{"c", "c++", "go", "python"}
	for i := 0; i < 4; i++ {
		fmt.Println(langs[i])
	}

	// namuna: with array 2
	for i, v := range langs {
		fmt.Println("Index:", i, "Value", v)
	}

	// namuna: with map
	var words = make(map[string]string)
	words["hello"] = "salom"
	words["how are you"] = "qalesan"
	for k, v := range words {
		fmt.Println("Key:", k, "Value:", v)
	}
}
