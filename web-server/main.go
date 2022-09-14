package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.fileServer(http.Dir("./static"))
	http.handle("/", fileServer)
	http.handleFunc("/hello", helloHandler)
	http.handleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8000\n")
	if err := http.listenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
