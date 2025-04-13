package product

import "gorm.io/gorm"

type ProductRepository interface {
	CreateProduct(product *Product) error
	GetAllProducts() ([]Product, error)
	GetProductByID(id uint) (*Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) CreateProduct(p *Product) error {
	return r.db.Create(p).Error
}

func (r *productRepository) GetAllProducts() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) GetProductByID(id uint) (*Product, error) {
	var p Product
	if err := r.db.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}
