package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello,</h1> you've requested: %s\n", r.URL.Path)
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	http.HandleFunc("/index", index)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8090", nil)
}
