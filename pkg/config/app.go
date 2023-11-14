// Package config provides functions to connect to a MySQL database using GORM.

package config

import "github.com/jinzhu/gorm"

// db holds the reference to the GORM database instance.
var db *gorm.DB

// Connect establishes a connection to the MySQL database.
func Connect() {
	// Open a new connection to the MySQL database
	d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		// Panic if there is an error connecting to the database
		panic(err)
	}
	// Set the global db variable to the opened database instance
	db = d
}

// GetDB returns the reference to the GORM database instance.
func GetDB() *gorm.DB {
	return db
}
