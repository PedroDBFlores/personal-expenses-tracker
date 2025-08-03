package ports

import (
	"pedro/personal-expenses-tracker/adapters/web/payloads"
	"pedro/personal-expenses-tracker/models"
)

type CreateExpensePort interface {
	CreateExpense(req payloads.CreateExpenseRequest) (*models.Expense, error)
}
