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
}