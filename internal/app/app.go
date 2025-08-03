package app

import (
	"go-subscription-service/internal/config"
	"go-subscription-service/internal/database"
	"go-subscription-service/internal/di"
	"go-subscription-service/internal/handler"
	"go-subscription-service/internal/migrations"

	"gorm.io/gorm"
)

type AppContainer struct {
	Config              *config.Config
	DB                  *gorm.DB
	SubscriptionHandler *handler.HTTPHandler
}

func InitApp() *AppContainer {
	cfg := config.Load()
	db := database.NewGorm(cfg.DatabaseDSN)

	migrations.Run(db)

	return &AppContainer{
		Config:              cfg,
		DB:                  db,
		SubscriptionHandler: di.InitSubscriptionService(db),
	}
}
