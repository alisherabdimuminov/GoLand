package main

import (
	"fmt"
	"io"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method, "URL:", r.URL)
	io.WriteString(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		panic(err)
	}
}
