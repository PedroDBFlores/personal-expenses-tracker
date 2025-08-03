package handlers

import (
	"strconv"

	"pedro/personal-expenses-tracker/ports"
	"pedro/personal-expenses-tracker/utils"

	"github.com/gofiber/fiber/v2"
)

type GetExpenseByIDHandler struct {
	Getter ports.GetExpensesPort
}

func NewGetExpenseByIDHandler(getter ports.GetExpensesPort) *GetExpenseByIDHandler {
	return &GetExpenseByIDHandler{Getter: getter}
}

func (h *GetExpenseByIDHandler) Handle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid expense ID")
	}
	expense, err := h.Getter.GetExpenseByID(uint(id))
	if err != nil {
		return utils.HandleError(c, err)
	}
	return c.JSON(expense)
}
