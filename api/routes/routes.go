package routes

import (
	"pedro/personal-expenses-tracker/adapters/database"
	"pedro/personal-expenses-tracker/adapters/web/handlers"
	"pedro/personal-expenses-tracker/usecases"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
)

func Setup(app *fiber.App) {
	dialector := sqlite.Open("expenses.db")
	db, _ := database.ConnectDatabase(dialector)

	// Build usecases
	getExpensesUC := usecases.NewGetExpensesUseCase(db)
	createExpensesUC := usecases.NewCreateExpenseUseCase(db)
	searchExpensesUC := usecases.NewSearchExpensesUseCase(db)
	updateExpensesUC := usecases.NewUpdateExpenseUseCase(db)
	deleteExpensesUC := usecases.NewDeleteExpenseUseCase(db)

	// Build handlers
	getExpensesHandler := handlers.NewGetExpensesHandler(getExpensesUC)
	createExpensesHandler := handlers.NewCreateExpenseHandler(createExpensesUC)
	searchExpensesHandler := handlers.NewSearchExpensesHandler(searchExpensesUC)
	updateExpensesHandler := handlers.NewUpdateExpenseHandler(updateExpensesUC)
	deleteExpensesHandler := handlers.NewDeleteExpenseHandler(deleteExpensesUC)

	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	expenses := api.Group("/expenses")
	expenses.Get("/", getExpensesHandler.Handle)
	expenses.Post("/", createExpensesHandler.Handle)
	expenses.Get("/search", searchExpensesHandler.Handle)
	expenses.Put("/:id", updateExpensesHandler.Handle)
	expenses.Delete("/:id", deleteExpensesHandler.Handle)
}
