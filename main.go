package main

import (
	"MyGram/database"
	"MyGram/routers"
	"log"
)

func init() {
	// When the program starts, it will automatically connect to the database
	database.StartDB()
}

func main() {
	// Run the server
	err := routers.StartServer().Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
