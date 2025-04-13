package order

import "gorm.io/gorm"

type OrderRepository interface {
	CreateOrder(order *Order) error
	GetOrdersByUser(userID uint) ([]Order, error)
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepo{db}
}

func (r *orderRepo) CreateOrder(order *Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepo) GetOrdersByUser(userID uint) ([]Order, error) {
	var orders []Order
	err := r.db.Preload("Items").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}
