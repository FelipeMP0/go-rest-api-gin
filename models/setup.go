package models

import (
	"github.com/jinzhu/gorm"
	// Required to initialize the database connection with sqlite
	_"github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

// DB represents the connection to the database instance.
var DB *gorm.DB

// ConnectDatabase creates a new connection with the database
// and initializes the datastore running the migrations.
func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "bookstore.db")

	if err != nil {
		log.Panic("Failed to connect to database")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
