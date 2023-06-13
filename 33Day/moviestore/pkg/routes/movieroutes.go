package routes

import (
	"github.com/Manda-supraja26/moviestore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterMovieRoutes = func(router *mux.Router) {
	router.HandleFun("/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFun("/movie/{bookid}", controllers.GetMovie).Methods("Get")
	router.HandleFun("/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFun("/movie/{bookid}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFun("/movie/{bookid}", controllers.DeleteMovie).Methods("DELETE")

}
