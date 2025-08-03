package main

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"pedro/personal-expenses-tracker/routes"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	app := fiber.New()
	routes.Setup(app)

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))

	logger.Info("API running", zap.String("addr", ":8080"))
	if err := app.Listen(":8080"); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
