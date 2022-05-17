package routes

import (
	"documentator/documentation/infraestructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func DocXasRoute(app *fiber.App, handler handlers.Handler) {
	users := app.Group("/api/documentation")

	users.Put("/update", handler.Edit)

	users.Get("/", handler.GetAll)

	users.Get("/filter/", handler.Find)

	users.Post("/new", handler.Create)
}
