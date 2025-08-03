package handlers

import (
	"pedro/personal-expenses-tracker/ports"
	"pedro/personal-expenses-tracker/utils"

	"github.com/gofiber/fiber/v2"
)

type GetExpensesHandler struct {
	Getter ports.GetExpensesPort
}

func NewGetExpensesHandler(getter ports.GetExpensesPort) *GetExpensesHandler {
	return &GetExpensesHandler{Getter: getter}
}

func (h *GetExpensesHandler) Handle(c *fiber.Ctx) error {
	expenses, err := h.Getter.GetAllExpenses()
	if err != nil {
		return utils.HandleError(c, err)
	}
	return c.JSON(expenses)
}
