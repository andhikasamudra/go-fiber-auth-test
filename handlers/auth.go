package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-auth-test/dto"
	"go-fiber-auth-test/services"
)

func Login(c *fiber.Ctx) error {
	var request dto.AuthLoginRequest

	err := c.BodyParser(request)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	response, err := services.Login(request)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	return c.JSON(response)
}
