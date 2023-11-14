// Package models defines the data models and database operations for the MovieHub API.

package models

import (
	"github.com/KarkiAnmol/MovieHub-API.git/pkg/config"
	"github.com/jinzhu/gorm"
)

// db holds the reference to the GORM database instance.
var db *gorm.DB

// Movie represents the model for a movie in the MovieHub.
type Movie struct {
	gorm.Model
	Name string `gorm:""json:"name"`
}

// init initializes the database connection and auto-migration for the Movie model.
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Movie{})
}

// GetMovie retrieves all movies from the database.
func GetMovie() []Movie {
	var ListOfMovies []Movie
	db.Find(&ListOfMovies)
	return ListOfMovies
}

// GetMovieById retrieves a movie by its ID from the database.
func GetMovieById(Id int64) (*Movie, *gorm.DB) {
	var MovieByID Movie
	db := db.Where("ID=?", Id).Find(&MovieByID)
	return &MovieByID, db
}

// CreateMovie creates a new movie record in the database.
func (m *Movie) CreateMovie() *Movie {
	db.NewRecord(m)
	db.Create(&m)
	return m
}

// DeleteMovie deletes a movie by its ID from the database.
func DeleteMovie(ID int64) Movie {
	var movie Movie
	db.Where("ID=?", ID).Delete(movie)
	return movie
}
