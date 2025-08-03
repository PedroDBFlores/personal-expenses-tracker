package payloads

import "time"

type CreateExpenseRequest struct {
	Amount            float64   `json:"amount"`
	Description       string    `json:"description"`
	Date              time.Time `json:"date"`
	ExpenseType       string    `json:"expenseType"`
	FulfillsExpenseId *uint     `json:"fulfillsExpenseId,omitempty"`
}
