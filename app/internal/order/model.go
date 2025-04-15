package order

import (
	"time"
)

type Order struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	CustomerID uint        `json:"customer_id"`
	ProductID  uint        `json:"product_id"`
	OrderDate  time.Time   `json:"order_date"`
	Amount     float64     `json:"amount"`
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
