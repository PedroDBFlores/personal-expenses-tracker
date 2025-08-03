package handlers

import (
	"pedro/personal-expenses-tracker/adapters/web/payloads"
	"pedro/personal-expenses-tracker/ports"
	"pedro/personal-expenses-tracker/utils"

	"github.com/gofiber/fiber/v2"
)

type SearchExpensesHandler struct {
	Searcher ports.SearchExpensesPort
}

func NewSearchExpensesHandler(searcher ports.SearchExpensesPort) *SearchExpensesHandler {
	return &SearchExpensesHandler{Searcher: searcher}
}

func (h *SearchExpensesHandler) Handle(c *fiber.Ctx) error {
	var req payloads.SearchExpensesRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	expenses, err := h.Searcher.SearchExpenses(req)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return c.JSON(expenses)
}
