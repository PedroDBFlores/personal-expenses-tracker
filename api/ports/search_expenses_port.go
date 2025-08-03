package ports

import (
	"pedro/personal-expenses-tracker/adapters/web/payloads"
	"pedro/personal-expenses-tracker/models"
)

type SearchExpensesPort interface {
	SearchExpenses(req payloads.SearchExpensesRequest) ([]models.Expense, error)
}
