package database

import (
	"log"
	"os"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var connection *gorm.DB

func ConnectDB() *gorm.DB {
	once.Do(func() {
		connection = initialize()
	})

	log.Println("Running Migrations...")

	return connection
}

func initialize() *gorm.DB {
	switch os.Getenv("DATABASE_DRIVER") {
	case "postgres":
		return initializePostgres()
	default:
		log.Fatalln("You must specify a database driver. Choices are 'postgres' or 'mysql'")
		return nil
	}
}
