package routes

import (
	"github.com/KarkiAnmol/MovieHub-API/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterMovieHubRoutes = func(router *mux.Router) {
	router.HandleFunc("/movie/", controllers.GetMovie).Methods("GET")           //get all movies
	router.HandleFunc("/movie/", controllers.CreateMovie).Methods("POST")       //create a new movie
	router.HandleFunc("/movie/{id}", controllers.GetMovieByID).Methods("GET")   //get one movie
	router.HandleFunc("/movie/{id}", controllers.UpdateMovie).Methods("PUT")    // update movie
	router.HandleFunc("/movie/{id}", controllers.DeleteMovie).Methods("DELETE") // Delete movie

}
