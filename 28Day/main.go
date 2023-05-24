package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID        string  `json:"id"`
	MovieName string  `json:"moviename"`
	Rating    float64 `json:"rating"`
	Language  string  `json:"language"`
	Director  string  `json:"director"`
}

var movie []Movie

func GetAllMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movie)
}

func GetMovieById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range movie {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
		}

	}

}

func DeleteMovieById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, v := range movie {
		if v.ID == params["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			break
		}
	}

}

func CreateNewMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var m Movie
	json.NewDecoder(req.Body).Decode(&m)
	rand.Seed(1000)
	m.ID = strconv.Itoa(rand.Intn(100000))
	movie = append(movie, m)
	json.NewEncoder(res).Encode(m)

}

// var *k
// *k
// &

func UpdateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for index, items := range movie {
		if items.ID == params["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			var m Movie
			json.NewDecoder(req.Body).Decode(&m)
			m.ID = params["id"]
			movie = append(movie, m)
			json.NewEncoder(res).Encode(movie)

		}

	}

}

func main() {
	movie = append(movie, Movie{ID: "1", MovieName: "RRR", Rating: 9.8, Language: "Telugu", Director: "Raja Moiuli"}, Movie{ID: "2", MovieName: "The kerala Story", Rating: 9.8, Language: "Malayalam", Director: "XXX"})

	r := mux.NewRouter()
	r.HandleFunc("/movies", GetAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovieById).Methods("Get")
	r.HandleFunc("/delete/{id}", DeleteMovieById).Methods("Delete")
	r.HandleFunc("/create", CreateNewMovie).Methods("POST")
	r.HandleFunc("/update/{id}", UpdateMovie).Methods("POST")

	fmt.Println("The server starting at port 8080")
	http.ListenAndServe(":8080", r)

}
