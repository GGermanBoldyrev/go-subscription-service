package service

import (
	"context"
	"github.com/google/uuid"
	"go-subscription-service/internal/model"
	"go-subscription-service/internal/repository"
	"time"
)

type SubscriptionService struct {
	repo *repository.SubscriptionRepository
}

func NewSubscriptionService(repo *repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) Create(ctx context.Context, sub *model.Subscription) error {
	return s.repo.Create(ctx, sub)
}

func (s *SubscriptionService) GetByID(ctx context.Context, id uuid.UUID) (*model.Subscription, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *SubscriptionService) Update(ctx context.Context, sub *model.Subscription) error {
	return s.repo.Update(ctx, sub)
}

func (s *SubscriptionService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *SubscriptionService) List(ctx context.Context) ([]*model.Subscription, error) {
	return s.repo.List(ctx)
}

func (s *SubscriptionService) TotalPrice(
	ctx context.Context,
	userID *uuid.UUID,
	serviceName string,
	from, to time.Time,
) (int, error) {
	return s.repo.TotalPrice(ctx, userID, serviceName, from, to)
}
