package report

import (
	"gorm.io/gorm"
)

type ReportRepository interface {
	GetTopCustomers() ([]TopCustomer, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db}
}

func (r *reportRepository) GetTopCustomers() ([]TopCustomer, error) {
	var topCustomers []TopCustomer

	query := `
			SELECT customer_id, SUM(amount) as total_spent
			FROM orders
			WHERE order_date >= NOW() - INTERVAL '1 month'
			GROUP BY customer_id
			ORDER BY total_spent DESC
			LIMIT 5
		`

	if err := r.db.Raw(query).Scan(&topCustomers).Error; err != nil {
		return nil, err
	}
	return topCustomers, nil
}
