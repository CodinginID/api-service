package order

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/CodinginID/api-service/internal/cart"
	"github.com/CodinginID/api-service/internal/product"
)

func RegisterOrderRoutes(r fiber.Router, db *gorm.DB) {
	// Inisialisasi dependency
	orderRepo := NewOrderRepository(db)
	cartRepo := cart.NewCartRepository(db)
	productRepo := product.NewProductRepository(db)

	orderService := NewOrderService(orderRepo, cartRepo, productRepo)
	orderHandler := NewOrderHandler(orderService)

	// Routes
	r.Post("/checkout", orderHandler.Checkout)
	r.Get("/history", orderHandler.GetOrderHistory)
}
