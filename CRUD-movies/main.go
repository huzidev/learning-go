package main

import (
	// 	"fmt"
	// 	"log"
	// 	"encoding/json"
	// 	"math/rand"
	// 	"net/http"
	// 	"strconv" // because id will INT and we've to make it String
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` // every movie has a director and all the types of Director
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	// functions
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("PUT")
	r.HandleFunc("/movies{id}", updateMovie).Methods("POST")
	r.HandleFunc("/movies{id}", deleteMovie).Methods("DELETE")
}
