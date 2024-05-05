package repositories

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"example.com/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetLocalCredentials retrieves database credentials from environment variables and returns them as a DatabaseCredentialsStore struct.
func GetLocalDbCredentials() models.DatabaseCreadentials {
	var info models.DatabaseCreadentials

	// Retrieve database connection information from environment variables
	info.Host = os.Getenv("host")             // Get the host address
	i, err := strconv.Atoi(os.Getenv("port")) // Get the port string and convert it into an int
	if err != nil {
		log.Fatal("Port should be int")
	}
	info.Port = i                         // Assign the port to the database connection information
	info.Database = os.Getenv("database") // Get the name of database
	info.Password = os.Getenv("password") // Get the password for database connection
	info.User = os.Getenv("user")         // Get the username for database connection

	// Return the database connection information
	return info
}

// ConnectDb Establish the database connection
func ConnectDb(local bool) (*gorm.DB, error) {
	var dbCreds models.DatabaseCreadentials

	// check if connection is local
	if local {
		dbCreds = GetLocalDbCredentials() // Retrieve database credentials
	} else {
		log.Println("Database connection is pending") // if not local
	}

	// Establish a connection to the database using the retrieved credentials
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbCreds.Host, dbCreds.Port, dbCreds.User, dbCreds.Password, dbCreds.Database, "disable")), &gorm.Config{})

	// Check if there was an error occur while establishing the connection
	if err != nil {
		log.Println(err)
		return db, err
	}

	// If there was no error then return the db and err(nil) successsfully
	return db, err
}
