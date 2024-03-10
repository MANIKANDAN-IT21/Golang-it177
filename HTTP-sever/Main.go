package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Manikandan!")
	})

	http.HandleFunc("/department", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Manikandan!")
	})

	http.HandleFunc("/roll", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, I'm 7376212it177!")
	})

	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Green")
	})

	fmt.Printf("Server running (port=8080), routes:\n")
	fmt.Printf("http://localhost:8080/hello\n")
	fmt.Printf("http://localhost:8080/name\n")
	fmt.Printf("http://localhost:8080/roll\n")
	fmt.Printf("http://localhost:8080/department\n")
	fmt.Printf("http://localhost:8080/color\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
