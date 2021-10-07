package routes

import (
	"path/controller"

	"github.com/gofiber/fiber/v2"
)

func CallRoutes(app fiber.Router) {
	app.Get("/", controller.GetAllData)
	app.Post("/", controller.CreateData)
	app.Get("/:id", controller.GetOneData)
	app.Delete("/:id", controller.DeleteOneData)
	app.Put("/:id", controller.UpdateData)

}
