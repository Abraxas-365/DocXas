package main

import (
	"documentator/documentation/application"
	"documentator/documentation/core/service"
	"documentator/documentation/infraestructure/repository"
	"documentator/documentation/infraestructure/rest/handlers"
	"documentator/documentation/infraestructure/rest/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	mongoUri := os.Getenv("MONGODB_URI")
	repo, err := repository.NewMongoRepository(mongoUri, "DocXas", 10, "Documentation")
	if err != nil {
		log.Panicln(err)
	}
	service := service.NewService(repo)

	application := application.NewApp(repo, service)

	handler := handlers.NewHandler(application)
	app := fiber.New()
	app.Use(logger.New())
	routes.DocXasRoute(app, handler)
	app.Listen(":3000")
}
