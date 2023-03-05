package main

import (
	"fmt"
	"net/http"
)

func greet() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
}
func main() {
	http.Handle("/", greet())
	http.ListenAndServe(":8080", nil)
}
