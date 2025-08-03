package payloads

import "time"

type SearchExpensesRequest struct {
	FromDate    *time.Time `json:"fromDate,omitempty"`
	ToDate      *time.Time `json:"toDate,omitempty"`
	MinAmount   *float64   `json:"minAmount,omitempty"`
	MaxAmount   *float64   `json:"maxAmount,omitempty"`
	ExpenseType *string    `json:"expenseType,omitempty"`
}
