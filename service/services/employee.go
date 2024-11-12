package services

import (
	"backend/models"
	database "backend/st_database"
	"context"
)

type EmployeeService struct {
	store database.IStore
}

func NewEmployeeService(store database.IStore) *EmployeeService {
	return &EmployeeService{
		store: store,
	}
}

func (s *EmployeeService) Create(ctx context.Context, req *models.Employee) (string, error) {
	id, err := s.store.Employee().Create(ctx, req)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (s *EmployeeService) Update(ctx context.Context, req *models.Employee) error {
	err := s.store.Employee().Update(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeService) Delete(ctx context.Context, req models.RequestId) error {
	err := s.store.Employee().Delete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeService) Get(ctx context.Context, req models.RequestId) (*models.Employee, error) {
	employee, err := s.store.Employee().Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (s *EmployeeService) GetAll(ctx context.Context, req models.GetAllEmployeesReq) (*models.GetAllEmployeesResp, error) {
	resp, err := s.store.Employee().GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
