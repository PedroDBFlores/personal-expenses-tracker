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

func setupTestDBForSearch(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.Expense{})
	assert.NoError(t, err)
	return db
}

func TestSearchExpensesUseCase_Success(t *testing.T) {
	db := setupTestDBForSearch(t)
	uc := usecases.NewSearchExpensesUseCase(db)

	// Create test data
	date1 := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC)
	date3 := time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC)

	expense1 := &models.Expense{Amount: 10.0, Description: "Food", ExpenseType: "Food", Date: date1}
	expense2 := &models.Expense{Amount: 20.0, Description: "Transport", ExpenseType: "Transport", Date: date2}
	expense3 := &models.Expense{Amount: 30.0, Description: "Food2", ExpenseType: "Food", Date: date3}

	db.Create(expense1)
	db.Create(expense2)
	db.Create(expense3)

	// Test filtering by amount
	minAmount := 15.0
	req := payloads.SearchExpensesRequest{MinAmount: &minAmount}

	results, err := uc.SearchExpenses(req)
	assert.NoError(t, err)
	assert.Len(t, results, 2) // expense2 and expense3

	// Test filtering by expense type
	expenseType := "Food"
	req2 := payloads.SearchExpensesRequest{ExpenseType: &expenseType}

	results2, err := uc.SearchExpenses(req2)
	assert.NoError(t, err)
	assert.Len(t, results2, 2) // expense1 and expense3

	// Test filtering by date range
	fromDate := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2025, 2, 15, 0, 0, 0, 0, time.UTC)
	req3 := payloads.SearchExpensesRequest{FromDate: &fromDate, ToDate: &toDate}

	results3, err := uc.SearchExpenses(req3)
	assert.NoError(t, err)
	assert.Len(t, results3, 1) // only expense2
}

func TestSearchExpensesUseCase_NoFilters(t *testing.T) {
	db := setupTestDBForSearch(t)
	uc := usecases.NewSearchExpensesUseCase(db)

	// Create test data
	expense1 := &models.Expense{Amount: 10.0, Description: "Food", ExpenseType: "Food", Date: time.Now()}
	expense2 := &models.Expense{Amount: 20.0, Description: "Transport", ExpenseType: "Transport", Date: time.Now()}
	db.Create(expense1)
	db.Create(expense2)

	// Search without any filters
	req := payloads.SearchExpensesRequest{}
	results, err := uc.SearchExpenses(req)
	assert.NoError(t, err)
	assert.Len(t, results, 2) // Should return all expenses
}

func TestSearchExpensesUseCase_EmptyResult(t *testing.T) {
	db := setupTestDBForSearch(t)
	uc := usecases.NewSearchExpensesUseCase(db)

	// Create test data
	expense1 := &models.Expense{Amount: 10.0, Description: "Food", ExpenseType: "Food", Date: time.Now()}
	db.Create(expense1)

	// Search with filters that should return no results
	minAmount := 100.0
	req := payloads.SearchExpensesRequest{MinAmount: &minAmount}

	results, err := uc.SearchExpenses(req)
	assert.NoError(t, err)
	assert.Len(t, results, 0) // Should return no results
}
