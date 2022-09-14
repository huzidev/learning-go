package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.fileServer(http.Dir("./static"))
}