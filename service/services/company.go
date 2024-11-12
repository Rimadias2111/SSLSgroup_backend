package services

import (
	"backend/models"
	database "backend/st_database"
	"context"
)

type CompanyService struct {
	store database.IStore
}

func NewCompanyService(store database.IStore) *CompanyService {
	return &CompanyService{store}
}

func (s *CompanyService) Create(ctx context.Context, company *models.Company) (string, error) {
	id, err := s.store.Company().Create(ctx, company)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *CompanyService) Update(ctx context.Context, company *models.Company) error {
	return s.store.Company().Update(ctx, company)
}

func (s *CompanyService) Delete(ctx context.Context, req models.RequestId) error {
	return s.store.Company().Delete(ctx, req)
}

func (s *CompanyService) Get(ctx context.Context, req models.RequestId) (*models.Company, error) {
	company, err := s.store.Company().Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (s *CompanyService) GetAll(ctx context.Context, req models.GetAllCompaniesReq) (*models.GetAllCompaniesResp, error) {
	resp, err := s.store.Company().GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
