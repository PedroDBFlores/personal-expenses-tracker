package ports

import (
	"pedro/personal-expenses-tracker/adapters/web/payloads"
	"pedro/personal-expenses-tracker/models"
)

type UpdateExpensePort interface {
	UpdateExpense(id uint, req payloads.UpdateExpenseRequest) (*models.Expense, error)
}
