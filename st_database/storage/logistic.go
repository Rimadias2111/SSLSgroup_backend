package storage

import (
	"backend/etc/helpers"
	"backend/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strconv"
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
	err := query.WithContext(ctx).Model(update).Where("id = ?", update.Id).
		Omit("Id", "DriverId").Updates(map[string]interface{}{
		"Post":       update.Post,
		"Status":     update.Status,
		"UpdateTime": update.UpdateTime,
		"StTime":     update.StTime,
		"State":      update.State,
		"Location":   update.Location,
		"Emoji":      update.Emoji,
		"Notion":     update.Notion,
		"CargoId":    update.CargoId,
	}).Error
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
		resp       models.GetAllLogisticsResp
		logistics  []models.LogisticResponse
		query      = s.db.WithContext(ctx).Model(&models.Logistic{}).Joins("JOIN drivers ON drivers.id = logistics.driver_id")
		offset     = (req.Page - 1) * req.Limit
		companyIds []uuid.UUID
	)

	if req.Status != "" {
		query = query.Where("logistics.status = ?", req.Status)
	}

	if req.Location != "" {
		query = query.Where("logistics.location = ?", req.Location)
	}

	if req.Type != "" {
		query = query.Where("drivers.type = ?", req.Type)
	}

	if req.Position != "" {
		query = query.Where("drivers.position = ?", req.Position)
	}

	if req.State != "" {
		query = query.Where("logistics.state = ?", req.State)
	}

	if req.Name != "" {
		query = query.Where("drivers.name = ?", req.Name)
	}

	if req.Post != "" {
		post, err := strconv.ParseBool(req.Post)
		if err != nil {
			return nil, err
		}
		query = query.Where("logistics.post = ?", post)
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
					logistics.updated_at as updated_at,
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
				WHEN 'READY AT HOME' THEN 3
				WHEN 'COVERED' THEN 4
				WHEN 'AT PU'  THEN 5
				WHEN 'ETA' THEN 6
				WHEN 'AT DEL' THEN 7
				WHEN 'ETA WILL BE LATE' THEN 8
				WHEN 'TRUCK ISSUES' THEN 9
				WHEN 'CANCELLED' THEN 10
				WHEN 'AT HOME' THEN 11
				WHEN 'LET US KNOW' THEN 12
				ELSE 999
			END
			`).
		Offset(int(offset)).
		Limit(int(req.Limit)).
		Scan(&logistics).
		Error
	if err != nil {
		return nil, errors.New("cannot get all logistics")
	}

	var companyOrder []uuid.UUID
	companyMap := map[uuid.UUID]*models.ByCompany{}
	for _, logistic := range logistics {
		if _, exists := companyMap[logistic.CompanyId]; !exists {
			companyMap[logistic.CompanyId] = &models.ByCompany{
				CompanyId:   logistic.CompanyId,
				CompanyName: "",
				Logistics:   []models.LogisticResponse{},
			}
			companyIds = append(companyIds, logistic.CompanyId)
			companyOrder = append(companyOrder, logistic.CompanyId)
		}
		companyMap[logistic.CompanyId].Logistics = append(companyMap[logistic.CompanyId].Logistics, logistic)
	}

	var companies []models.Company
	err = s.db.WithContext(ctx).Model(&models.Company{}).Where("id IN (?)", companyIds).Select(`id, CONCAT(name, '   ', scac) AS name`).Scan(&companies).Error
	if err != nil {
		return nil, err
	}

	for _, company := range companies {
		if companyInfo, exists := companyMap[company.Id]; exists {
			companyInfo.CompanyName = company.Name
		}
	}

	resp.Companies = make([]models.ByCompany, 0, len(companyOrder))
	for _, companyId := range companyOrder {
		resp.Companies = append(resp.Companies, *companyMap[companyId])
	}

	helpers.CountDown(&resp)

	countQuery := s.db.WithContext(ctx).Model(&models.Logistic{}).Joins("JOIN drivers ON drivers.id = logistics.driver_id")
	if req.Name != "" {
		countQuery = countQuery.Where("drivers.name ILIKE ?", "%"+req.Name+"%")
	}

	if req.Type != "" {
		countQuery = countQuery.Where("drivers.type = ?", req.Type)
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

func (s *LogisticRepo) Overview(ctx context.Context) (models.GetOverview, error) {
	var (
		resp  models.GetOverview
		query = s.db.WithContext(ctx).Model(&models.Logistic{}).Joins("JOIN drivers ON drivers.id = logistics.driver_id")
	)

	err := query.
		Select(`
            '' AS name,
            drivers.company_id AS id,
            COUNT(CASE 
                WHEN logistics.status IN ('READY', 'READY AT HOME') THEN 1 
                END) AS free_drivers,
            COUNT(CASE 
                WHEN (logistics.status IN ('ETA', 'ETA, WILL BE LATE')
											AND logistics.st_time <= NOW() + INTERVAL '1 hour') 
                     OR logistics.status IN ('WILL BE READY', 'AT DEL') THEN 1 
                END) AS will_be_soon_drivers,
            COUNT(CASE 
                WHEN logistics.status IN ('COVERED', 'AT PU') OR (logistics.status IN ('ETA', 'ETA, WILL BE LATE')
																		AND logistics.st_time >= NOW() + INTERVAL '1 hour') THEN 1 
                END) AS occupied_drivers,
            COUNT(CASE
                WHEN logistics.status IN ('LET US KNOW', 'AT HOME') THEN 1
                END) AS not_working
        `).
		Group("drivers.company_id").
		Scan(&resp.Companies).Error

	if err != nil {
		return resp, err
	}

	return resp, nil
}
