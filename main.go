package main

import "MyGram/database"

func init() {
	// When the program starts, it will automatically connect to the database
	database.StartDB()
}

func main() {

}
