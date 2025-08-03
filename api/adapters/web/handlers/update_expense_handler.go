package handlers

import (
	"strconv"

	"pedro/personal-expenses-tracker/adapters/web/payloads"
	"pedro/personal-expenses-tracker/ports"
	"pedro/personal-expenses-tracker/utils"

	"github.com/gofiber/fiber/v2"
)

type UpdateExpenseHandler struct {
	Updater ports.UpdateExpensePort
}

func NewUpdateExpenseHandler(updater ports.UpdateExpensePort) *UpdateExpenseHandler {
	return &UpdateExpenseHandler{Updater: updater}
}

func (h *UpdateExpenseHandler) Handle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid expense ID")
	}
	var req payloads.UpdateExpenseRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	expense, err := h.Updater.UpdateExpense(uint(id), req)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return c.JSON(expense)
}
