package cart

import "time"

type CartService interface {
	Add(userID, productID uint, qty int) error
	List(userID uint) ([]CartItem, error)
	Clear(userID uint) error
}

type cartService struct {
	repo CartRepository
}


func NewCartService(r CartRepository) CartService {
	return &cartService{r}
}

func (s *cartService) Add(userID, productID uint, qty int) error {
	item := &CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  qty,
		CreatedAt: time.Now(),
	}
	return s.repo.AddItem(item)
}

func (s *cartService) List(userID uint) ([]CartItem, error) {
	return s.repo.GetItemsByUser(userID)
}

func (s *cartService) Clear(userID uint) error {
	return s.repo.ClearCart(userID)
}
