// Package main handles the database setup
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DeskDistance handles the post data structure and the database model.
type DeskDistance struct {
	gorm.Model
	Distance uint `json:"distance"`
}

func getDb() *gorm.DB {

	// Create or get the database.
	db, err := gorm.Open(sqlite.Open("deskDistance.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&DeskDistance{})

	return db
}
