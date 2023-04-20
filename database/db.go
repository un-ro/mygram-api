package database

import (
	"MyGram/models"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func initDB() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	USERNAME := os.Getenv("DB_USERNAME")
	PASSWORD := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", HOST, USERNAME, PASSWORD, PORT, DBNAME)
}

func StartDB() {
	dsn := initDB()

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Invalid connect to database")
	}

	log.Println("Success connect to database")

	err := db.Migrator().DropTable(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	if err != nil {
		log.Fatalln("Error drop table")
	}

	err = db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	if err != nil {
		log.Fatalln("Error migrate database")
	}
}

func GetDB() *gorm.DB {
	return db
}
