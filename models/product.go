package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Brand            string    `json:"brand"`
	Price            float64   `json:"price"`
	Discount         float64   `json:"discount"`
	DiscountedPrice  float64   `json:"discounted_price"`
	InstallmentValue float64   `json:"installment_value"`
	MaxInstallments  int       `json:"max_installments"`
	HasInterest      bool      `json:"has_interest"`
	HighlightStatus  *string   `json:"highlight_status,omitempty"`
	StoreName        string    `json:"store_name"`
	ImageLink        string    `json:"image_link"`
	Category         string    `json:"category"`
	IsDeleted        bool      `json:"is_deleted"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
