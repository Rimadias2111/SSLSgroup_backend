package service

import "backend/service/services"

type IService interface {
	Company() *services.CompanyService
	Driver() *services.DriverService
	Employee() *services.EmployeeService
}

type Service struct {
	companyService  *services.CompanyService
	driverService   *services.DriverService
	employeeService *services.EmployeeService
}
