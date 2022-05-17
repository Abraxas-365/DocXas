package handlers

import "github.com/gofiber/fiber/v2"

func (handler *handler) GetAll(c *fiber.Ctx) error {
	documentatios, err := handler.app.GetAll()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(documentatios)
}
