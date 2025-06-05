package models

import "time"

type Order struct {
	ID        uint        `gorm:"primaryKey"`
	UserID    uint        `gorm:"not null"`
	Items     []OrderItem `gorm:"foreignKey:OrderID"`
	Total     float64     `gorm:"not null"`
	Status    string      `gorm:"default:'pending'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
