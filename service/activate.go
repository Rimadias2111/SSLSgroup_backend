package service

import (
	"backend/service/services"
	database "backend/st_database"
)

func New(store database.IStore) IService {
	return &Service{
		companyService:     services.NewCompanyService(store),
		driverService:      services.NewDriverService(store),
		employeeService:    services.NewEmployeeService(store),
		logisticService:    services.NewLogisticService(store),
		transactionService: services.NewTransactionService(store),
		performanceService: services.NewPerformanceService(store),
	}
}

func (s *Service) Company() *services.CompanyService { return s.companyService }

func (s *Service) Driver() *services.DriverService { return s.driverService }

func (s *Service) Employee() *services.EmployeeService { return s.employeeService }

func (s *Service) Logistic() *services.LogisticService { return s.logisticService }

func (s *Service) Transaction() *services.TransactionService { return s.transactionService }

func (s *Service) Performance() *services.PerformanceService { return s.performanceService }
