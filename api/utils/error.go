package utils

import (
	"github.com/gofiber/fiber/v2"
)

// HandleError writes a JSON error response with status 500 (or 400 for fiber errors)
func HandleError(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if fiberErr, ok := err.(*fiber.Error); ok {
		code = fiberErr.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
