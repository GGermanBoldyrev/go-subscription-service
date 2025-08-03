package di

import (
	"go-subscription-service/internal/handler"
	"go-subscription-service/internal/repository"
	"go-subscription-service/internal/service"
	"gorm.io/gorm"
)

func InitSubscriptionService(db *gorm.DB) *handler.HTTPHandler {
	subRepo := repository.NewSubscriptionRepository(db)
	subSvc := service.NewSubscriptionService(subRepo)
	subHandler := handler.NewHTTPHandler(subSvc)

	return subHandler
}
