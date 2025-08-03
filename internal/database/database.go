package database

import (
	"go-subscription-service/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}
