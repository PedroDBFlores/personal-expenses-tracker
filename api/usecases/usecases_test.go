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

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.Expense{})
	assert.NoError(t, err)
	return db
}

func TestCreateExpenseUseCase(t *testing.T) {
	db := setupTestDB(t)
	uc := usecases.NewCreateExpenseUseCase(db)
	req := payloads.CreateExpenseRequest{
		Amount:      20.0,
		Description: "Test",
		Date:        time.Now(),
		ExpenseType: "TestType",
	}
	exp, err := uc.CreateExpense(req)
	assert.NoError(t, err)
	assert.Equal(t, req.Amount, exp.Amount)
	assert.Equal(t, req.Description, exp.Description)
}

// Similar tests can be written for update, delete, get, search use cases using the in-memory DB.

func TestGetExpensesUseCase_GetAllExpenses(t *testing.T) {
	db := setupTestDB(t)
	uc := usecases.NewGetExpensesUseCase(db)

	// Create test data
	expense1 := &models.Expense{Amount: 10.0, Description: "Test1", ExpenseType: "Food", Date: time.Now()}
	expense2 := &models.Expense{Amount: 20.0, Description: "Test2", ExpenseType: "Transport", Date: time.Now()}
	db.Create(expense1)
	db.Create(expense2)

	expenses, err := uc.GetAllExpenses()
	assert.NoError(t, err)
	assert.Len(t, expenses, 2)
	assert.Equal(t, "Test1", expenses[0].Description)
	assert.Equal(t, "Test2", expenses[1].Description)
}

func TestGetExpensesUseCase_GetExpenseByID(t *testing.T) {
	db := setupTestDB(t)
	uc := usecases.NewGetExpensesUseCase(db)

	// Create test data
	expense := &models.Expense{Amount: 15.0, Description: "Test Expense", ExpenseType: "Food", Date: time.Now()}
	db.Create(expense)

	retrieved, err := uc.GetExpenseByID(expense.ID)
	assert.NoError(t, err)
	assert.NotNil(t, retrieved)
	assert.Equal(t, expense.ID, retrieved.ID)
	assert.Equal(t, expense.Description, retrieved.Description)
}

func TestGetExpensesUseCase_GetExpenseByID_NotFound(t *testing.T) {
	db := setupTestDB(t)
	uc := usecases.NewGetExpensesUseCase(db)

	_, err := uc.GetExpenseByID(999)
	assert.Error(t, err)
}

// Tests extracted to individual use case test files:
// - create_expense_test.go (CreateExpenseUseCase tests)
// - get_expenses_test.go (GetExpensesUseCase tests)
// - update_expense_test.go (UpdateExpenseUseCase tests)
// - delete_expense_test.go (DeleteExpenseUseCase tests)
// - search_expenses_test.go (SearchExpensesUseCase tests)
