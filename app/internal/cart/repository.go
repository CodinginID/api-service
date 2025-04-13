package cart

import "gorm.io/gorm"

type CartRepository interface {
	AddItem(item *CartItem) error
	GetItemsByUser(userID uint) ([]CartItem, error)
	ClearCart(userID uint) error
}

type cartRepo struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepo{db}
}

func (r *cartRepo) AddItem(item *CartItem) error {
	return r.db.Create(item).Error
}

func (r *cartRepo) GetItemsByUser(userID uint) ([]CartItem, error) {
	var items []CartItem
	err := r.db.Where("user_id = ?", userID).Find(&items).Error
	return items, err
}

func (r *cartRepo) ClearCart(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&CartItem{}).Error
}
