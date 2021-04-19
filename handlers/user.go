package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-auth-test/dto"
	"go-fiber-auth-test/services"
)

func CreateUser(c *fiber.Ctx) error {
	var request dto.CreateUserRequest

	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	_, err = services.CreateUser(request)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "User Created",
	})
}
