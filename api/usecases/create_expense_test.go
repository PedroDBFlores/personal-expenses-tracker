package usecases_test

import (
	"pedro/personal-expenses-tracker/adapters/web/payloads"
	"pedro/personal-expenses-tracker/models"
	"pedro/personal-expenses-tracker/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDBForCreate(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.Expense{})
	assert.NoError(t, err)
	return db
}

func TestCreateExpenseUseCase_ValidExpense(t *testing.T) {
	db := setupTestDBForCreate(t)
	uc := usecases.NewCreateExpenseUseCase(db)

	req := payloads.CreateExpenseRequest{
		Amount:      25.50,
		Description: "Grocery shopping",
		Date:        time.Date(2025, 8, 10, 14, 30, 0, 0, time.UTC),
		ExpenseType: "Food",
	}

	expense, err := uc.CreateExpense(req)

	assert.NoError(t, err)
	assert.NotNil(t, expense)
	assert.NotZero(t, expense.ID)
	assert.Equal(t, req.Amount, expense.Amount)
	assert.Equal(t, req.Description, expense.Description)
	assert.Equal(t, req.Date, expense.Date)
	assert.Equal(t, req.ExpenseType, expense.ExpenseType)
	assert.NotZero(t, expense.CreatedAt)
	assert.NotZero(t, expense.UpdatedAt)
}

func TestCreateExpenseUseCase_WithFulfillsExpenseId(t *testing.T) {
	db := setupTestDBForCreate(t)
	uc := usecases.NewCreateExpenseUseCase(db)

	// Create a base expense first
	baseExpense := &models.Expense{
		Amount:      100.0,
		Description: "Base expense",
		Date:        time.Now(),
		ExpenseType: "Debit",
	}
	db.Create(baseExpense)

	// Create expense that fulfills the base expense
	req := payloads.CreateExpenseRequest{
		Amount:            100.0,
		Description:       "Payment for base expense",
		Date:              time.Now(),
		ExpenseType:       "Credit",
		FulfillsExpenseId: &baseExpense.ID,
	}

	expense, err := uc.CreateExpense(req)

	assert.NoError(t, err)
	assert.NotNil(t, expense)
	assert.Equal(t, req.FulfillsExpenseId, expense.FulfillsExpenseId)
	assert.Equal(t, baseExpense.ID, *expense.FulfillsExpenseId)
}

func TestCreateExpenseUseCase_MinimalData(t *testing.T) {
	db := setupTestDBForCreate(t)
	uc := usecases.NewCreateExpenseUseCase(db)

	req := payloads.CreateExpenseRequest{
		Amount:      1.0,
		Description: "Test",
		Date:        time.Now(),
		ExpenseType: "Test",
	}

	expense, err := uc.CreateExpense(req)

	assert.NoError(t, err)
	assert.NotNil(t, expense)
	assert.Equal(t, req.Amount, expense.Amount)
	assert.Equal(t, req.Description, expense.Description)
	assert.Equal(t, req.ExpenseType, expense.ExpenseType)
	assert.Nil(t, expense.FulfillsExpenseId)
}
