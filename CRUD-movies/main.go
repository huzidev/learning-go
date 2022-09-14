package main

import (
	// 	"fmt"
	// 	"log"
	// 	"encoding/json"
	// 	"math/rand"
	// 	"net/http"
	// 	"strconv" // because id will INT and we've to make it String
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// encode complete movies slice which is defined above var movies []Movie
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	// by default movies
	movies = append(movies,
		Movie{
			ID:    "1",
			Isbn:  "123456",
			Title: "Dark",
			Director: &Director{
				FirstName: "Sohaib",
				LastName:  "Hassan",
			},
		},
	)
	movies = append(movies,
		Movie{
			ID:    "2",
			Isbn:  "654321",
			Title: "Dragon ball",
			Director: &Director{
				FirstName: "Huzi",
				LastName:  "Boss",
			},
		},
	)

	// functions
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	port := 8000
	fmt.Printf("Starting server at port %v\n", port)
	log.Fatal(http.ListenAndServe(":8000", r))
}
