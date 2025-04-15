package db

import (
	"log"

	"github.com/CodinginID/api-service/internal/auth"
	"github.com/CodinginID/api-service/internal/cart"
	"github.com/CodinginID/api-service/internal/order"
	"github.com/CodinginID/api-service/internal/product"
	"gorm.io/gorm"
)

func RunAutoMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&auth.User{},
		&product.Product{},
		&cart.CartItem{},
		&order.Order{},
		&order.OrderItem{},
	)
	if err != nil {
		log.Fatalf("AutoMigration failed: %v", err)
	}

	log.Println("✅ Database migration completed.")
}
