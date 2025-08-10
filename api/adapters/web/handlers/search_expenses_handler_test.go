package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pedro/personal-expenses-tracker/adapters/web/handlers"
	"pedro/personal-expenses-tracker/adapters/web/payloads"
	"pedro/personal-expenses-tracker/models"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockSearchExpensesPortForSearch struct{ mock.Mock }

func (m *mockSearchExpensesPortForSearch) SearchExpenses(req payloads.SearchExpensesRequest) ([]models.Expense, error) {
	args := m.Called(req)
	return args.Get(0).([]models.Expense), args.Error(1)
}

func TestSearchExpensesHandler_Success(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockSearchExpensesPortForSearch)
	handler := handlers.NewSearchExpensesHandler(mockPort)
	app.Post("/expenses/search", handler.Handle)

	reqBody := `{"minAmount":5.0}`
	expenses := []models.Expense{
		{ID: 1, Amount: 10.5, Description: "Lunch", ExpenseType: "Food"},
	}
	mockPort.On("SearchExpenses", mock.MatchedBy(func(req payloads.SearchExpensesRequest) bool {
		return req.MinAmount != nil && *req.MinAmount == 5.0
	})).Return(expenses, nil)

	req := httptest.NewRequest(http.MethodPost, "/expenses/search", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var got []models.Expense
	_ = json.NewDecoder(resp.Body).Decode(&got)
	assert.Len(t, got, 1)
	assert.Equal(t, expenses[0].ID, got[0].ID)
	mockPort.AssertExpectations(t)
}

func TestSearchExpensesHandler_InvalidJSON(t *testing.T) {
	app := fiber.New()
	mockPort := new(mockSearchExpensesPortForSearch)
	handler := handlers.NewSearchExpensesHandler(mockPort)
	app.Post("/expenses/search", handler.Handle)

	reqBody := `{"invalid": json}`
	req := httptest.NewRequest(http.MethodPost, "/expenses/search", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
