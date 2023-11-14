// Package routes defines the API routes for the MovieHub API.

package routes

import (
	"github.com/KarkiAnmol/MovieHub-API.git/pkg/controllers"
	"github.com/gorilla/mux"
)

// RegisterMovieHubRoutes registers all the routes for the MovieHub API.
var RegisterMovieHubRoutes = func(router *mux.Router) {
	router.HandleFunc("/movie/", controllers.GetMovie).Methods("GET")           // Get all movies
	router.HandleFunc("/movie/", controllers.CreateMovie).Methods("POST")       // Create a new movie
	router.HandleFunc("/movie/{id}", controllers.GetMovieById).Methods("GET")   // Get one movie
	router.HandleFunc("/movie/{id}", controllers.UpdateMovie).Methods("PUT")    // Update movie
	router.HandleFunc("/movie/{id}", controllers.DeleteMovie).Methods("DELETE") // Delete movie
}
