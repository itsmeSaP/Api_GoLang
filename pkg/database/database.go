package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	//syntax to connect to the database go_practice created by us
	url := "postgres://postgres:sherlock16@localhost:5432/go_practice"
	connection, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	db = connection
	return db, err
}

func GetDB() *gorm.DB {
	return db
}
