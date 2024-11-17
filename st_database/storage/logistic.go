package storage

import (
	"backend/etc/helpers"
	"backend/models"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogisticRepo struct {
	db *gorm.DB
}

func NewLogisticRepo(db *gorm.DB) Logistic {
	return &LogisticRepo{
		db: db,
	}
}

func (s *LogisticRepo) Create(ctx context.Context, update *models.Logistic, tx ...*gorm.DB) (string, error) {
	var (
		id    = uuid.New()
		query = s.db
	)
	update.Id = id

	if len(tx) > 0 && tx[0] != nil {
		query = tx[0]
	}

	err := query.WithContext(ctx).Create(&update).Error
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *LogisticRepo) Update(ctx context.Context, update *models.Logistic, tx ...*gorm.DB) error {
	var query = s.db
	if len(tx) > 0 && tx[0] != nil {
		query = tx[0]
	}
	err := query.WithContext(ctx).Model(update).
		Omit("Id", "DriverId").Updates(update).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *LogisticRepo) Delete(ctx context.Context, req models.RequestId) error {
	err := s.db.WithContext(ctx).Model(&models.Logistic{}).Where("id = ?", req.Id).Delete(&models.Logistic{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *LogisticRepo) Get(ctx context.Context, req models.RequestId) (*models.Logistic, error) {
	var update models.Logistic

	err := s.db.WithContext(ctx).Model(&models.Logistic{}).Preload("Driver").Preload("Cargo").Where("id = ?", req.Id).First(&update).Error
	if err != nil {
		return nil, err
	}

	return &update, nil
}

func (s *LogisticRepo) GetAll(ctx context.Context, req models.GetAllLogisticsReq) (*models.GetAllLogisticsResp, error) {
	var (
		resp   models.GetAllLogisticsResp
		query  = s.db.WithContext(ctx).Model(&models.Logistic{}).Joins("JOIN drivers ON drivers.id = logistics.driver_id")
		offset = (req.Page - 1) * req.Limit
	)

	if req.Name != "" {
		query = query.Where("drivers.name ILIKE ?", "%"+req.Name+"%")
	}

	if req.Status != "" {
		query = query.Where("logistics.status = ?", req.Status)
	}

	if req.Location != "" {
		query = query.Where("logistics.location = ?", req.Location)
	}

	if req.Type != "" {
		query = query.Where("drivers.type = ?", req.Type)
	}

	err := query.Select(`
					logistics.id as id,
					logistics.post as post,
					logistics.driver_id as driver_id,
					logistics.status as status,
					logistics.update_time as update_time,
					logistics.st_time as st_time,
					logistics.state as state,
					logistics.location as location,
					logistics.notion as notion,
					logistics.emoji as emoji,
					logistics.cargo_id as cargo_id,
					drivers.name as driver_name,
					drivers.surname as driver_surname,
					drivers.type as driver_type,
					drivers.position as driver_position,
					drivers.company_id as company_id
					`).
		Order("drivers.company_id ASC").
		Order(`
             CASE logistics.status
				WHEN 'READY' THEN 1
				WHEN 'WILL BE READY' THEN 2
				WHEN 'COVERED' THEN 3
				WHEN 'AT PU'  THEN 4
				WHEN 'ETA' THEN 5
				WHEN 'AT DEL' THEN 6
				WHEN 'ETA WILL BE LATE' THEN 7
				WHEN 'TRUCK ISSUES' THEN 8
				WHEN 'CANCELLED' THEN 9
				WHEN 'AT HOME' THEN 10
				WHEN 'LET US KNOW' THEN 11
				ELSE 999
			END
			`).
		Offset(int(offset)).
		Limit(int(req.Limit)).
		Scan(&resp.Logistics).
		Error
	if err != nil {
		return nil, err
	}

	var companyIds []uuid.UUID
	var companies []models.Company
	for _, body := range resp.Logistics {
		companyIds = append(companyIds, body.CompanyId)
	}

	err = s.db.WithContext(ctx).Model(&models.Company{}).Where("id IN (?)", companyIds).Find(&companies).Error
	if err != nil {
		return nil, err
	}

	companyMap := make(map[uuid.UUID]string)
	for _, company := range companies {
		companyMap[company.Id] = company.Name
	}

	for i := range resp.Logistics {
		if name, exists := companyMap[resp.Logistics[i].CompanyId]; exists {
			resp.Logistics[i].CompanyName = name
		}
	}

	helpers.CountDown(&resp)

	countQuery := s.db.WithContext(ctx).Model(&models.Logistic{})
	if req.Name != "" {
		countQuery = countQuery.Joins("Driver").
			Where("drivers.name ILIKE ?", "%"+req.Name+"%")

		if req.Type != "" {
			countQuery = countQuery.Where("drivers.type = ?", req.Type)
		}
	}
	if req.Status != "" {
		countQuery = countQuery.Where("logistics.status = ?", req.Status)
	}
	if req.Location != "" {
		countQuery = countQuery.Where("logistics.location = ?", req.Location)
	}

	err = countQuery.Count(&resp.Count).Error
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
