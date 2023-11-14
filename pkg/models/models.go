package models

import (
	"github.com/KarkiAnmol/MovieHub-API.git/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Movie struct {
	gorm.Model
	Name string `gorm:""json:"name"`
}

//initialize connection to database and auto-migration for the Movie model
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Movie{})
}
func GetMovie() []Movie {
	var ListOfMovies []Movie
	db.Find(&ListOfMovies)
	return ListOfMovies
}
func GetMovieById(Id int64) (*Movie, *gorm.DB) {
	var MovieByID Movie
	db := db.Where("ID=?", Id).Find(&MovieByID)
	return &MovieByID, db
}
func (m *Movie) CreateMovie() *Movie {
	db.NewRecord(m)
	db.Create(&m)
	return m
}

//the returned Movie struct will be parsed into JSON to be displayed to the user in Postman
func DeleteMovie(ID int64) Movie {
	var movie Movie
	db.Where("ID=?", ID).Delete(movie)
	return movie
}
