package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sheexiong/Go-REST-API/bin/models"
)

// InitialMigration for project with db.AutoMigrate
func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&models.Property{}) {
		db.CreateTable(&models.Property{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Property{})
	}

	if !db.HasTable(&models.Country{}) {
		db.CreateTable(&models.Country{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Country{})
	}

	return db
}
