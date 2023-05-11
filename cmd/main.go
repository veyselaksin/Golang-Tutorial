package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/veyselaksin/Golang-Tutorial/database"
	"github.com/veyselaksin/Golang-Tutorial/models"
	"github.com/veyselaksin/Golang-Tutorial/routes"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App, conn *gorm.DB) {

	userService := routes.NewUser(conn)

	// User Endpoints
	app.Post("/api/users", userService.CreateUser)
	app.Get("/api/users", userService.GetUsers)
}

func main() {
	connection := database.ConnectDB()

	sqlDB, err := connection.DB()
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// TODO: Add migrations
	connection.AutoMigrate(
		&models.User{},
		// &models.Product{},
		// &models.Order{},
	)

	app := fiber.New()

	setupRoutes(app, connection)

	log.Fatal(app.Listen(":8000"))
}
