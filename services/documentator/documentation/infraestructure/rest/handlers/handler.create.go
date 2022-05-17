package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) Create(c *fiber.Ctx) error {
	body := struct {
		Url string `json:"url"`
	}{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	fmt.Println(body.Url)
	if err := handler.app.Create(body.Url); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
