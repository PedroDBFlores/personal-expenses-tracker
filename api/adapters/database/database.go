package database

import (
	"pedro/personal-expenses-tracker/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func ConnectDatabase(dialector gorm.Dialector) (*gorm.DB, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var err error
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database", zap.Error(err))
	}
	if err := db.AutoMigrate(&models.Expense{}); err != nil {
		logger.Fatal("failed to migrate database", zap.Error(err))
	}
	return db, nil
}
