package handlers

import (
	"documentator/documentation/application"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Create(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	Find(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}

type handler struct {
	app application.Application
}

func NewHandler(app application.Application) Handler {
	return &handler{
		app,
	}
}
