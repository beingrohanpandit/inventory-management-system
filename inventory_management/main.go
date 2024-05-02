package main

import (
	"log"
	"os"
	"strconv"

	"example.com/inventory_management/handlers"
	"example.com/inventory_management/repositories"
	"github.com/gin-gonic/gin"
)

func init() {

	// Get the local (true or false) and convert it into bool value
	local, _ := strconv.ParseBool(os.Getenv("local"))

	// Initialize the gin router
	r := gin.Default()

	// Connect to the database
	db, err := repositories.ConnectDb(local)
	if err != nil {
		log.Println(err)
	}

	// Initialize an instance of the application with the database connection
	app := handlers.MyApp{
		DB: db,
	}

	// Define routes for the application
	app.Routes(r)

	// If running locally then start the server
	if local {
		r.Run(":8080")
	}

}

func main() {

}
