package model

import (
	"github.com/google/uuid"
	"time"
)

type Subscription struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	ServiceName string     `gorm:"not null" json:"service_name" example:"Yandex Plus"`
	Price       int        `gorm:"not null" json:"price" example:"400"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null" json:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   time.Time  `gorm:"not null" json:"start_date" example:"2025-07-01T00:00:00Z"`
	EndDate     *time.Time `json:"end_date,omitempty" example:"2025-12-01T00:00:00Z"`
}
