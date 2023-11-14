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
