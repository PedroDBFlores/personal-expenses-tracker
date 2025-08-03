package payloads

import "time"

type UpdateExpenseRequest struct {
	Amount            *float64   `json:"amount,omitempty"`
	Description       *string    `json:"description,omitempty"`
	Date              *time.Time `json:"date,omitempty"`
	ExpenseType       *string    `json:"expenseType,omitempty"`
	FulfillsExpenseId *uint      `json:"fulfillsExpenseId,omitempty"`
}
