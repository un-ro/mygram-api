package database

import (
	"MyGram/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	PORT     = "32768"
	USERNAME = "postgres"
	PASSWORD = "postgrespw"
	DBNAME   = "fga"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", HOST, USERNAME, PASSWORD, PORT, DBNAME)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Invalid connect to database")
	}

	log.Println("Success connect to database")

	err := db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	if err != nil {
		log.Fatalln("Error migrate database")
	}
}

func GetDB() *gorm.DB {
	return db
}
