package handlers

import (
	"strconv"

	"pedro/personal-expenses-tracker/ports"
	"pedro/personal-expenses-tracker/utils"

	"github.com/gofiber/fiber/v2"
)

type DeleteExpenseHandler struct {
	Deleter ports.DeleteExpensePort
}

func NewDeleteExpenseHandler(deleter ports.DeleteExpensePort) *DeleteExpenseHandler {
	return &DeleteExpenseHandler{Deleter: deleter}
}

func (h *DeleteExpenseHandler) Handle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid expense ID")
	}
	if err := h.Deleter.DeleteExpense(uint(id)); err != nil {
		return utils.HandleError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
