package handlers

import (
	"pedro/personal-expenses-tracker/adapters/web/payloads"
	"pedro/personal-expenses-tracker/ports"
	"pedro/personal-expenses-tracker/utils"

	"github.com/gofiber/fiber/v2"
)

type CreateExpenseHandler struct {
	Creator ports.CreateExpensePort
}

func NewCreateExpenseHandler(creator ports.CreateExpensePort) *CreateExpenseHandler {
	return &CreateExpenseHandler{Creator: creator}
}

func (h *CreateExpenseHandler) Handle(c *fiber.Ctx) error {
	var req payloads.CreateExpenseRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	expense, err := h.Creator.CreateExpense(req)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return c.Status(fiber.StatusCreated).JSON(expense)
}
