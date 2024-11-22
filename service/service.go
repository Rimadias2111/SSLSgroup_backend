package service

import "backend/service/services"

type IService interface {
	Company() *services.CompanyService
	Driver() *services.DriverService
	Employee() *services.EmployeeService
	Logistic() *services.LogisticService
	Transaction() *services.TransactionService
	Performance() *services.PerformanceService
}

type Service struct {
	companyService     *services.CompanyService
	driverService      *services.DriverService
	employeeService    *services.EmployeeService
	logisticService    *services.LogisticService
	transactionService *services.TransactionService
	performanceService *services.PerformanceService
}
