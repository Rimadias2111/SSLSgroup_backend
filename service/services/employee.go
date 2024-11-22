package services

import (
	"backend/etc/helpers"
	"backend/etc/jwt"
	"backend/models"
	database "backend/st_database"
	"context"
	"errors"
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

func (s *EmployeeService) Auth(ctx context.Context, req models.AuthReq) (models.AuthResp, error) {
	var resp models.AuthResp
	employee, err := s.store.Employee().GetByUsername(ctx, req.Username)
	if err != nil {
		return resp, errors.New("did not find employee")
	}

	matched, err := helpers.CheckPassword(req.Password, employee.Password)
	if err != nil {
		return resp, err
	}
	if !matched {
		return resp, errors.New("password not match")
	}

	resp.Token, err = jwt.GenerateToken(employee.Id.String(), employee.Username, int(employee.AccessLevel))
	if err != nil {
		return resp, err
	}
	resp.Employee = employee
	return resp, nil
}
