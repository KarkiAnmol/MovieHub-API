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

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	createMovie := &models.Movie{}
	utils.ParseBody(r, createMovie)
	b := createMovie.CreateMovie()
	res, _ := json.Marshal(b) //Marshals the created movie (b) into JSON format. This prepares the movie data to be included in the HTTP response.
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["ID"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	deleteMovie := models.DeleteMovie(ID)
	res, _ := json.Marshal(deleteMovie)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetMovie(w http.ResponseWriter, r *http.Request) {
	movieDetails := models.GetMovie()
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["Id"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	movieDetails, _ := models.GetMovieById(ID)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["ID"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	updateMovie := &models.Movie{}
	utils.ParseBody(r, updateMovie)
	movieDetails, db := models.GetMovieById(ID)
	if updateMovie.Name != "" {
		movieDetails.Name = updateMovie.Name
	}
	db.Save(&movieDetails)
	res, _ := json.Marshal(updateMovie)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
