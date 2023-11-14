package main

import (
	"log"
	"net/http"

	"github.com/KarkiAnmol/MovieHub-API.git/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Create a new Gorilla mux router
	r := mux.NewRouter()

	// Register routes for the MovieHub API
	routes.RegisterMovieHubRoutes(r)

	// Handle all incoming requests using the created router
	http.Handle("/", r)

	// Start the HTTP server and listen on localhost:9010
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
