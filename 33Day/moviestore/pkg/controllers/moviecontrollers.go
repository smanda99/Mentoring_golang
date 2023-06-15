package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Manda-supraja26/moviestore/pkg/models"
	"github.com/Manda-supraja26/moviestore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewMovie models.Movie

// New Movie
func CreateMovie(res http.ResponseWriter, req *http.Request) {
	createMovie := &models.Movie{}
	utils.ParseBody(req, createMovie)
	movie := createMovie.CreateMovie()
	r, _ := json.Marshal(movie)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)
}

// GetAll Movies
func GetAllMovies(res http.ResponseWriter, req *http.Request) {
	newMovie := models.GetAllMovies()
	r, _ := json.Marshal(newMovie)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)
}

// Get Movie By ID
func GetMovie(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	movieId := vars["movieid"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error with parsing")
		// log.Fatal(err)
	}
	moviedetails, _ := models.GetMovieByID(ID)
	r, _ := json.Marshal(moviedetails)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)
}

// Delete Movie
func DeleteMovie(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	movieID := vars["movieid"]
	ID, err := strconv.ParseInt(movieID, 0, 0)
	if err != nil {
		fmt.Println("error with parsing and converting ")
	}
	movie := models.DeleteMovie(ID)
	r, _ := json.Marshal(movie)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)
}

func UpdateMovie(res http.ResponseWriter, req *http.Request) {
	var createMovie = &models.Movie{}
	utils.ParseBody(req, createMovie)
	vars := mux.Vars(req)
	movieID := vars["movieid"]
	ID, err := strconv.ParseInt(movieID, 0, 0)
	if err != nil {
		fmt.Println("error with parsing and converting ")
	}

	// var createMovie = &models.Movie{}
	// utils.Parse(req, createMovie)

	bookdetails, db := models.GetMovieByID(ID)

	if createMovie.Name != "" {
		bookdetails.Name = createMovie.Name
	}
	if createMovie.Title != "" {
		bookdetails.Title = createMovie.Title
	}
	if createMovie.Year != 0 {
		bookdetails.Year = createMovie.Year
	}
	if createMovie.Rating != 0.0 {
		bookdetails.Rating = createMovie.Rating
	}
	db.Save(&bookdetails)
	r, _ := json.Marshal(bookdetails)
	res.Header().Set("Content-TYpe", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r)

}
