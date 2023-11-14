package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/KarkiAnmol/MovieHub-API.git/pkg/models"
	"github.com/KarkiAnmol/MovieHub-API.git/pkg/utils"
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
