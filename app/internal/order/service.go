package order

import (
	"errors"
	"time"

	"github.com/CodinginID/api-service/internal/cart"
	"github.com/CodinginID/api-service/internal/product"
)

type OrderService interface {
	Checkout(userID uint) (*Order, error)
	GetOrderHistory(userID uint) ([]Order, error)
}

type orderService struct {
	orderRepo   OrderRepository
	cartRepo    cart.CartRepository
	productRepo product.ProductRepository
}

func NewOrderService(or OrderRepository, cr cart.CartRepository, pr product.ProductRepository) OrderService {
	return &orderService{
		orderRepo:   or,
		cartRepo:    cr,
		productRepo: pr,
	}
}

func (s *orderService) Checkout(userID uint) (*Order, error) {
	// 1. Get cart items
	items, err := s.cartRepo.GetItemsByUser(userID)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, errors.New("cart is empty")
	}

	var orderItems []OrderItem
	var total float64

	// 2. Build OrderItem from cart
	for _, item := range items {
		product, err := s.productRepo.GetProductByID(item.ProductID)
		if err != nil {
			return nil, err
		}

		subtotal := float64(item.Quantity) * product.Price
		total += subtotal

		orderItems = append(orderItems, OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})
	}

	// 3. Save order
	order := &Order{
		CustomerID: userID,
		ProductID:  orderItems[0].ProductID, // Assuming the first product in the order
		OrderDate:  time.Now(),
		Amount:     total,
		CreatedAt:  time.Now(),
		Items:      orderItems,
	}

	if err := s.orderRepo.CreateOrder(order); err != nil {
		return nil, err
	}

	// 4. Clear cart
	if err := s.cartRepo.ClearCart(userID); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderService) GetOrderHistory(userID uint) ([]Order, error) {
	return s.orderRepo.GetOrdersByUser(userID)
}
