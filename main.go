package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
	"log"
	"path/database"
	"path/routes"
)

func handler(app *fiber.App) {
	
	routes.CallRoutes(app)
}

func main() {

	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	database.ConnectDB()

	handler(app)

	app.Listen(":8080")

}
