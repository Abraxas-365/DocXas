package handlers

import (
	"documentator/documentation/core/models"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) Edit(c *fiber.Ctx) error {
	edit := new(models.Documentation)
	if err := c.BodyParser(&edit); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if err := handler.app.Edit(*edit); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
