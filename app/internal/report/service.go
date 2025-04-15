package report

type ReportService interface {
	GetTopCustomers() ([]TopCustomer, error)
}

type reportService struct {
	repo ReportRepository
}

func NewReportService ( repo ReportRepository) ReportService {
	return &reportService{repo}
}
func (s *reportService) GetTopCustomers() ([]TopCustomer, error) {
	topCustomers, err := s.repo.GetTopCustomers()
	if err != nil {
		return nil, err
	}
	return topCustomers, nil
}