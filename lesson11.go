// usr/bin/go
// author: Ali
// date  : 27/09/2023
package main

import (
	"fmt"
)

func main() {
	// Go dasturlash tilida nested maps (ichma ich maps)

	var users = make(map[int]map[string]string)

	users[1] = make(map[string]string)
	users[1]["first_name"] = "Lololo"
	users[1]["last_name"] = "Pepepe"

	users[2] = make(map[string]string)
	users[2]["first_name"] = "Pepepe"
	users[2]["last_name"] = "Lololo"

	fmt.Println(users)

	for key, user := range users {
		fmt.Println(key)
		for k, v := range user {
			fmt.Println("\t", k, v)
		}
	}
}
