package routes

import (
	"pedro/personal-expenses-tracker/adapters/database"
	"pedro/personal-expenses-tracker/adapters/web/handlers"
	"pedro/personal-expenses-tracker/ports"
	"pedro/personal-expenses-tracker/usecases"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Setup(app *fiber.App) {
	dialector := sqlite.Open("expenses.db")
	db, _ := database.ConnectDatabase(dialector)

	api := app.Group("/expenses")
	getExpensesHandler := handlers.NewGetExpensesHandler(getExpenseGetter(db))
	createExpensesHandler := handlers.NewCreateExpenseHandler(getExpenseCreator(db))
	api.Get("/", getExpensesHandler.Handle)
	api.Post("/", createExpensesHandler.Handle)
}

func getExpenseGetter(db *gorm.DB) ports.GetExpensesPort {
	return usecases.NewGetExpensesUseCase(db)
}

func getExpenseCreator(db *gorm.DB) ports.CreateExpensePort {
	return usecases.NewCreateExpenseUseCase(db)
}
