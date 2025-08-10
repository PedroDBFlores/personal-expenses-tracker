package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"pedro/personal-expenses-tracker/adapters/web/handlers"
	"pedro/personal-expenses-tracker/adapters/web/payloads"
	"pedro/personal-expenses-tracker/models"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCreateExpensePortIndividual struct{ mock.Mock }

func (m *mockCreateExpensePortIndividual) CreateExpense(req payloads.CreateExpenseRequest) (*models.Expense, error) {
	args := m.Called(req)
	return args.Get(0).(*models.Expense), args.Error(1)
}

func TestCreateExpenseHandler_ValidRequest(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockCreateExpensePortIndividual)
	handler := handlers.NewCreateExpenseHandler(mockPort)
	app.Post("/expenses", handler.Handle)

	reqBody := `{"amount":12.50,"description":"Lunch","date":"2025-08-10T00:00:00Z","expenseType":"Food"}`
	expected := &models.Expense{
		ID:          1,
		Amount:      12.50,
		Description: "Lunch",
		Date:        time.Date(2025, 8, 10, 0, 0, 0, 0, time.UTC),
		ExpenseType: "Food",
	}

	mockPort.On("CreateExpense", mock.MatchedBy(func(req payloads.CreateExpenseRequest) bool {
		return req.Amount == 12.50 && req.Description == "Lunch"
	})).Return(expected, nil)

	req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response models.Expense
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expected.ID, response.ID)
	assert.Equal(t, expected.Amount, response.Amount)
	assert.Equal(t, expected.Description, response.Description)

	mockPort.AssertExpectations(t)
}

func TestCreateExpenseHandler_DatabaseError(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockCreateExpensePortIndividual)
	handler := handlers.NewCreateExpenseHandler(mockPort)
	app.Post("/expenses", handler.Handle)

	reqBody := `{"amount":12.50,"description":"Lunch","date":"2025-08-10T00:00:00Z","expenseType":"Food"}`
	mockPort.On("CreateExpense", mock.Anything).Return((*models.Expense)(nil), errors.New("database connection failed"))

	req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	mockPort.AssertExpectations(t)
}

func TestCreateExpenseHandler_InvalidRequestBody(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockCreateExpensePortIndividual)
	handler := handlers.NewCreateExpenseHandler(mockPort)
	app.Post("/expenses", handler.Handle)

	// Invalid JSON
	reqBody := `{"amount": "invalid", "description": "Lunch"}`
	req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateExpenseHandler_MissingContentType(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockCreateExpensePortIndividual)
	handler := handlers.NewCreateExpenseHandler(mockPort)
	app.Post("/expenses", handler.Handle)

	reqBody := `{"amount":12.50,"description":"Lunch","date":"2025-08-10T00:00:00Z","expenseType":"Food"}`
	req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(reqBody))
	// Missing Content-Type header
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateExpenseHandler_Success_FromMainFile(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockCreateExpensePortIndividual)
	handler := handlers.NewCreateExpenseHandler(mockPort)
	app.Post("/expenses", handler.Handle)

	reqBody := `{"amount":10.5,"description":"Lunch","date":"2025-08-10T00:00:00Z","expenseType":"Food"}`
	expected := &models.Expense{ID: 1, Amount: 10.5, Description: "Lunch", Date: time.Date(2025, 8, 10, 0, 0, 0, 0, time.UTC), ExpenseType: "Food"}
	mockPort.On("CreateExpense", mock.Anything).Return(expected, nil)

	req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	var got models.Expense
	_ = json.NewDecoder(resp.Body).Decode(&got)
	assert.Equal(t, expected.ID, got.ID)
	mockPort.AssertExpectations(t)
}
