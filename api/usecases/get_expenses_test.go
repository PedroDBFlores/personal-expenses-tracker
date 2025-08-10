package usecases_test

import (
	"pedro/personal-expenses-tracker/models"
	"pedro/personal-expenses-tracker/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDBForGet(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.Expense{})
	assert.NoError(t, err)
	return db
}

func TestGetExpensesUseCase_GetAllExpenses_MultipleRecords(t *testing.T) {
	db := setupTestDBForGet(t)
	uc := usecases.NewGetExpensesUseCase(db)

	// Create test data
	expenses := []*models.Expense{
		{Amount: 10.0, Description: "Coffee", ExpenseType: "Food", Date: time.Date(2025, 8, 1, 0, 0, 0, 0, time.UTC)},
		{Amount: 50.0, Description: "Groceries", ExpenseType: "Food", Date: time.Date(2025, 8, 2, 0, 0, 0, 0, time.UTC)},
		{Amount: 25.0, Description: "Gas", ExpenseType: "Transport", Date: time.Date(2025, 8, 3, 0, 0, 0, 0, time.UTC)},
	}

	for _, expense := range expenses {
		db.Create(expense)
	}

	result, err := uc.GetAllExpenses()

	assert.NoError(t, err)
	assert.Len(t, result, 3)
	assert.Equal(t, "Coffee", result[0].Description)
	assert.Equal(t, "Groceries", result[1].Description)
	assert.Equal(t, "Gas", result[2].Description)
}

func TestGetExpensesUseCase_GetAllExpenses_EmptyDatabase(t *testing.T) {
	db := setupTestDBForGet(t)
	uc := usecases.NewGetExpensesUseCase(db)

	result, err := uc.GetAllExpenses()

	assert.NoError(t, err)
	assert.Len(t, result, 0)
	assert.NotNil(t, result) // Should return empty slice, not nil
}

func TestGetExpensesUseCase_GetExpenseByID_ExistingRecord(t *testing.T) {
	db := setupTestDBForGet(t)
	uc := usecases.NewGetExpensesUseCase(db)

	// Create test data
	expense := &models.Expense{
		Amount:      75.25,
		Description: "Restaurant dinner",
		ExpenseType: "Food",
		Date:        time.Date(2025, 8, 10, 19, 30, 0, 0, time.UTC),
	}
	db.Create(expense)

	result, err := uc.GetExpenseByID(expense.ID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expense.ID, result.ID)
	assert.Equal(t, expense.Amount, result.Amount)
	assert.Equal(t, expense.Description, result.Description)
	assert.Equal(t, expense.ExpenseType, result.ExpenseType)
	assert.Equal(t, expense.Date, result.Date)
}

func TestGetExpensesUseCase_GetExpenseByID_NonExistentRecord(t *testing.T) {
	db := setupTestDBForGet(t)
	uc := usecases.NewGetExpensesUseCase(db)

	result, err := uc.GetExpenseByID(999)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "record not found")
}

func TestGetExpensesUseCase_GetExpenseByID_WithFulfillsExpenseId(t *testing.T) {
	db := setupTestDBForGet(t)
	uc := usecases.NewGetExpensesUseCase(db)

	// Create base expense
	baseExpense := &models.Expense{
		Amount:      100.0,
		Description: "Base expense",
		ExpenseType: "Debit",
		Date:        time.Now(),
	}
	db.Create(baseExpense)

	// Create fulfilling expense
	fulfillingExpense := &models.Expense{
		Amount:            100.0,
		Description:       "Payment",
		ExpenseType:       "Credit",
		Date:              time.Now(),
		FulfillsExpenseId: &baseExpense.ID,
	}
	db.Create(fulfillingExpense)

	result, err := uc.GetExpenseByID(fulfillingExpense.ID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.FulfillsExpenseId)
	assert.Equal(t, baseExpense.ID, *result.FulfillsExpenseId)
}
