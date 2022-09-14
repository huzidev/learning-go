package main

import (
	"fmt"
	"log"
	"net/http"
)

// response with w and request with r and * (pointer)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	// at "/hello" user will see (Hello, World) and it's a GET method
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello, World")
}

func formHandler(w http.ResponseWriter, r *http.Request)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
