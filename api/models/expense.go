package models

import (
	"time"
)

type Expense struct {
	ID                uint `gorm:"primaryKey;autoIncrement"`
	Amount            float64
	Description       string
	Date              time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ExpenseType       string
	FulfillsExpenseId *uint
}
