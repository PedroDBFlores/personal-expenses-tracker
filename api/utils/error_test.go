package utils

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHandleError_Default500(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return HandleError(c, errors.New("something went wrong"))
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestHandleError_FiberError(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return HandleError(c, fiber.ErrBadRequest)
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, 400, resp.StatusCode)
}
