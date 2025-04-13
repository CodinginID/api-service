package order

import (
	"time"
)

type Order struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	UserID     uint        `json:"user_id"`
	TotalPrice float64     `json:"total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	Items      []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
