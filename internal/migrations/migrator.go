package migrations

import (
	"go-subscription-service/internal/model"
	"go-subscription-service/pkg/logger"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Subscription{},
	)
	if err != nil {
		logger.Fatalf("migration failed: %v", err)
	}
}
