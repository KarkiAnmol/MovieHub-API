// Package controllers provides HTTP request handlers for MovieHub API.

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KarkiAnmol/MovieHub-API.git/pkg/models"
	"github.com/KarkiAnmol/MovieHub-API.git/pkg/utils"
	"github.com/gorilla/mux"
)

// CreateMovie handles the creation of a new movie record.
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON body of the request into a Movie struct
	createMovie := &models.Movie{}
	utils.ParseBody(r, createMovie)

	// Create a new movie record in the database
	b := createMovie.CreateMovie()

	// Marshal the created movie into JSON format for the HTTP response
	res, _ := json.Marshal(b)

	// Set the content type header to "application/json" and respond with the created movie details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteMovie handles the deletion of a movie record by ID.
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	// Extract movie ID from request parameters
	vars := mux.Vars(r)
	Id := vars["ID"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	// Delete the movie record from the database
	deleteMovie := models.DeleteMovie(ID)

	// Marshal the deleted movie details into JSON format for the HTTP response
	res, _ := json.Marshal(deleteMovie)

	// Set the content type header to "application/json" and respond with the deleted movie details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetMovie retrieves all movie records from the database.
func GetMovie(w http.ResponseWriter, r *http.Request) {
	// Retrieve all movie records from the database
	movieDetails := models.GetMovie()

	// Marshal the movie details into JSON format for the HTTP response
	res, _ := json.Marshal(movieDetails)

	// Set the content type header to "application/json" and respond with all movie details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetMovieById retrieves a specific movie record by ID from the database.
func GetMovieById(w http.ResponseWriter, r *http.Request) {
	// Extract movie ID from request parameters
	vars := mux.Vars(r)
	Id := vars["Id"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	// Retrieve the movie record by ID from the database
	movieDetails, _ := models.GetMovieById(ID)

	// Marshal the movie details into JSON format for the HTTP response
	res, _ := json.Marshal(movieDetails)

	// Set the content type header to "application/json" and respond with the specific movie details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateMovie handles the update of a movie record by ID.
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	// Extract movie ID from request parameters
	vars := mux.Vars(r)
	Id := vars["ID"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	// Parse the JSON body of the request into a Movie struct
	updateMovie := &models.Movie{}
	utils.ParseBody(r, updateMovie)

	// Retrieve the existing movie record by ID from the database
	movieDetails, db := models.GetMovieById(ID)

	// Update the movie details if provided in the request
	if updateMovie.Name != "" {
		movieDetails.Name = updateMovie.Name
	}

	// Save the updated movie details to the database
	db.Save(&movieDetails)

	// Marshal the updated movie details into JSON format for the HTTP response
	res, _ := json.Marshal(updateMovie)

	// Set the content type header to "application/json" and respond with the updated movie details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
