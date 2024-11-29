package storage

import (
	"backend/models"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strconv"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) Transaction {
	return &TransactionRepo{
		db: db,
	}
}

func (t *TransactionRepo) Create(ctx context.Context, transaction *models.Transaction, tx ...*gorm.DB) (string, error) {
	var (
		id    = uuid.New()
		query = t.db
	)
	transaction.Id = id
	if len(tx) > 0 && tx[0] != nil {
		query = tx[0]
	}

	if err := query.WithContext(ctx).Create(&transaction).Error; err != nil {
		return "", err
	}

	return id.String(), nil
}

func (t *TransactionRepo) Update(ctx context.Context, transaction *models.Transaction) error {
	return t.db.WithContext(ctx).Model(transaction).Omit("Id").Updates(transaction).Error
}

func (t *TransactionRepo) Delete(ctx context.Context, req models.RequestId) error {
	return t.db.WithContext(ctx).Where("id = ?", req).Delete(&models.Transaction{}).Error
}

func (t *TransactionRepo) Get(ctx context.Context, req models.RequestId) (*models.Transaction, error) {
	var transaction models.Transaction
	err := t.db.WithContext(ctx).Where("id = ?", req.Id).Preload("Driver").First(&transaction).Error
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (t *TransactionRepo) GetAll(ctx context.Context, req models.GetAllTransReq) (*models.GetAllTransResp, error) {
	var (
		resp   models.GetAllTransResp
		offset = (req.Page - 1) * req.Limit
		query  = t.db.WithContext(ctx).Model(&models.Transaction{}).Preload("Driver")
	)

	if req.CargoID != "" {
		query = query.Where("cargo_id = ?", req.CargoID)
	}

	if req.Provider != "" {
		query = query.Where("provider = ?", req.Provider)
	}

	if req.Success != "" {
		success, err := strconv.ParseBool(req.Success)
		if err != nil {
			return nil, err
		}
		query = query.Where("success = ?", success)
	}

	err := query.Find(&resp.Transactions).Offset(int(offset)).Limit(int(req.Page)).Error
	if err != nil {
		return nil, err
	}

	err = query.Count(&resp.Count).Error
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
