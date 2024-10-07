package models

import (
	"github.com/google/uuid"
	"time"
)

type ProductReview struct {
	ID           uuid.UUID `json:"id"`
	ProductID    uuid.UUID `json:"product_id"`
	TotalReviews int       `json:"total_reviews"`
	ReviewScore  float64   `json:"review_score"`
	IsDeleted    bool      `json:"is_deleted"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
