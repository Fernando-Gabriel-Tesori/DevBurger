package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	CustomerName string  `json:"customer_name"`
	Items        string  `json:"items"` // Pode ser JSON string dos itens
	TotalPrice   float64 `json:"total_price"`
	Status       string  `json:"status"`
}
