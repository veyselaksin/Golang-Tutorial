package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/veyselaksin/Golang-Tutorial/database"
	"github.com/veyselaksin/Golang-Tutorial/models"
	"github.com/veyselaksin/Golang-Tutorial/routes"
)

func setupRoutes(app *fiber.App) {

	// User Endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
}

func main() {
	connection := database.ConnectDB()

	// TODO: Add migrations
	connection.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
	)

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
