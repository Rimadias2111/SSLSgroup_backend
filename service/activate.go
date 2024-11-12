package service

import (
	"backend/service/services"
	database "backend/st_database"
)

func New(store database.IStore) Service {
	return Service{
		companyService:  services.NewCompanyService(store),
		driverService:   services.NewDriverService(store),
		employeeService: services.NewEmployeeService(store),
	}
}

func (s *Service) Company() *services.CompanyService { return s.companyService }

func (s *Service) Driver() *services.DriverService { return s.driverService }

func (s *Service) Employee() *services.EmployeeService { return s.employeeService }
