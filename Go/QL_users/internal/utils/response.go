package utils

import "github.com/gofiber/fiber/v2"

func Success(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"success": true,
			"data":    data,
		},
	)
}
func OK(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"success": true,
			"message": message,
		},
	)
}
func Created(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"success": true,
			"message": message,
		},
	)
}

func BadRequest(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(
		fiber.Map{
			"success": false,
			"message": message,
		},
	)
}

func Unauthorized(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(
		fiber.Map{
			"success": false,
			"message": message,
		},
	)
}

func Forbidden(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusForbidden).JSON(
		fiber.Map{
			"success": false,
			"message": message,
		},
	)
}

func NotFound(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(
		fiber.Map{
			"success": false,
			"message": message,
		},
	)
}

func InternalServerError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(
		fiber.Map{
			"success": false,
			"message": message,
		},
	)
}