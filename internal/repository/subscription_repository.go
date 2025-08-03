package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go-subscription-service/internal/model"
	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Create(ctx context.Context, sub *model.Subscription) error {
	return r.db.WithContext(ctx).Create(sub).Error
}

func (r *SubscriptionRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Subscription, error) {
	var sub model.Subscription
	if err := r.db.WithContext(ctx).First(&sub, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &sub, nil
}

func (r *SubscriptionRepository) Update(ctx context.Context, sub *model.Subscription) error {
	return r.db.WithContext(ctx).Save(sub).Error
}

func (r *SubscriptionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.Subscription{}, "id = ?", id).Error
}

func (r *SubscriptionRepository) List(ctx context.Context) ([]*model.Subscription, error) {
	var subs []*model.Subscription
	err := r.db.WithContext(ctx).Find(&subs).Error
	if err != nil {
		return nil, err
	}
	return subs, nil
}
