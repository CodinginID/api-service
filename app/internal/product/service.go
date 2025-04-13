package product

import "time"

type ProductService interface {
	Create(name, desc string, price float64, stock int) error
	GetAll() ([]Product, error)
	GetByID(id uint) (*Product, error)
}

type productService struct {
	repo ProductRepository
}

func NewProductService(r ProductRepository) ProductService {
	return &productService{r}
}

func (s *productService) Create(name, desc string, price float64, stock int) error {
	p := &Product{
		Name:        name,
		Description: desc,
		Price:       price,
		Stock:       stock,
		CreatedAt:   time.Now(),
	}
	return s.repo.CreateProduct(p)
}

func (s *productService) GetAll() ([]Product, error) {
	return s.repo.GetAllProducts()
}

func (s *productService) GetByID(id uint) (*Product, error) {
	return s.repo.GetProductByID(id)
}
