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

// @title My Gram API
// @version 1.0
// @description Final Project Hacktiv8: Scalable microservices Web API with Golang
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header -> Bearer
// @name Authorization
func main() {
	// Run the server
	err := routers.StartServer().Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
