package main

import (
	"log"
	"net/http"

	"github.com/Manda-supraja26/moviestore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterMovieRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))

}
