package migrations

import (
	"go-subscription-service/internal/model"
	"gorm.io/gorm"
	"log"
)

func Run(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Subscription{},
	)
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}
