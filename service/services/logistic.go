package services

import (
	"backend/etc/Utime"
	"backend/models"
	"backend/models/swag"
	database "backend/st_database"
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogisticService struct {
	store database.IStore
}

func NewLogisticService(store database.IStore) *LogisticService {
	return &LogisticService{store: store}
}

func (s *LogisticService) Create(ctx context.Context, req *models.Logistic) (string, error) {
	id, err := s.store.Logistic().Create(ctx, req)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *LogisticService) Update(ctx context.Context, req *models.Logistic) error {
	err := s.store.Logistic().Update(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *LogisticService) Delete(ctx context.Context, req models.RequestId) error {
	err := s.store.Logistic().Delete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *LogisticService) Get(ctx context.Context, req models.RequestId) (*models.Logistic, error) {
	resp, err := s.store.Logistic().Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *LogisticService) GetAll(ctx context.Context, req models.GetAllLogisticsReq) (*models.GetAllLogisticsResp, error) {
	resp, err := s.store.Logistic().GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *LogisticService) UpdateWithCargo(ctx context.Context, logistic *models.Logistic, cargo *models.Cargo, create bool) (string, error) {
	var (
		db  = s.store.DB()
		id  string
		err error
	)
	transErr := db.Transaction(func(tx *gorm.DB) error {
		if create {
			id, err = s.store.Cargo().Create(ctx, cargo, tx)
			if err != nil {
				return err
			}

			cargoId, errP := uuid.Parse(id)
			if errP != nil {
				return errP
			}

			logistic.CargoId = &cargoId
		} else {
			err = s.store.Cargo().Update(ctx, cargo, tx)
			if err != nil {
				return err
			}
			id = cargo.Id.String()
		}

		err = s.store.Logistic().Update(ctx, logistic, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if transErr != nil {
		return "", transErr
	}

	return id, nil
}

func (s *LogisticService) Terminate(ctx context.Context, req models.RequestId, success bool) error {
	db := s.store.DB()
	err := db.Transaction(func(tx *gorm.DB) error {
		logistic, err := s.store.Logistic().Get(ctx, req)
		if err != nil {
			return err
		}
		_, err = s.store.Transaction().Create(ctx, &models.Transaction{
			From:         logistic.Cargo.From,
			To:           logistic.Cargo.To,
			PuTime:       logistic.Cargo.PickUpTime,
			DeliveryTime: logistic.Cargo.DeliveryTime,
			LoadedMiles:  logistic.Cargo.LoadedMiles,
			TotalMiles:   logistic.Cargo.LoadedMiles + logistic.Cargo.FreeMiles,
			Provider:     logistic.Cargo.Provider,
			Cost:         logistic.Cargo.Cost,
			Rate:         logistic.Cargo.Rate,
			DriverId:     logistic.DriverId,
			EmployeeId:   logistic.Cargo.EmployeeId,
			CargoID:      logistic.Cargo.CargoID,
			Success:      success,
		}, tx)
		if err != nil {
			return err
		}

		errU := s.store.Logistic().Update(ctx, &models.Logistic{
			Id:         logistic.Id,
			Post:       false,
			Status:     "READY",
			UpdateTime: Utime.Now(),
			StTime:     nil,
			State:      logistic.State,
			Location:   logistic.Location,
			Emoji:      "",
			Notion:     "",
			CargoId:    nil,
		}, tx)
		if errU != nil {
			return errU
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *LogisticService) CancelLate(ctx context.Context, req swag.CancelLogistic, reqId models.RequestId, idEmp models.RequestId) error {
	db := s.store.DB()

	err := db.Transaction(func(tx *gorm.DB) error {
		logistic, getErr := s.store.Logistic().Get(ctx, reqId)
		if getErr != nil {
			return getErr
		}

		if logistic.CargoId == nil {
			return errors.New("cargo not found")
		}

		var employee *models.Employee = &models.Employee{}
		if idEmp.Id != uuid.Nil {
			employee, getErr = s.store.Employee().Get(ctx, idEmp)
			if getErr != nil {
				return getErr
			}
		} else {
			employee.Name = ""
			employee.Surname = ""
		}

		if req.Cancel {
			_, err := s.store.Performance().Create(ctx, &models.Performance{
				Reason:     req.Reason,
				WhoseFault: req.WhoseFault,
				Status:     req.Status,
				Section:    req.Section,
				DisputedBy: employee.Name + employee.Surname,
				Company:    req.Company,
				LoadId:     logistic.Cargo.CargoID,
			}, tx)
			if err != nil {
				return err
			}

			_, err = s.store.Transaction().Create(ctx, &models.Transaction{
				From:         logistic.Cargo.From,
				To:           logistic.Cargo.To,
				PuTime:       logistic.Cargo.PickUpTime,
				DeliveryTime: logistic.Cargo.DeliveryTime,
				LoadedMiles:  logistic.Cargo.LoadedMiles,
				TotalMiles:   logistic.Cargo.LoadedMiles + logistic.Cargo.FreeMiles,
				Provider:     logistic.Cargo.Provider,
				Cost:         logistic.Cargo.Cost,
				Rate:         logistic.Cargo.Rate,
				DriverId:     logistic.DriverId,
				EmployeeId:   logistic.Cargo.EmployeeId,
				CargoID:      logistic.Cargo.CargoID,
				Success:      false,
			}, tx)
			if err != nil {
				return err
			}

			errU := s.store.Logistic().Update(ctx, &models.Logistic{
				Id:         logistic.Id,
				Post:       false,
				Status:     "READY",
				UpdateTime: Utime.Now(),
				StTime:     nil,
				State:      logistic.State,
				Location:   logistic.Location,
				Emoji:      "",
				Notion:     "",
				CargoId:    nil,
			}, tx)
			if errU != nil {
				return errU
			}
		} else {
			_, err := s.store.Performance().Create(ctx, &models.Performance{
				Reason:     req.Reason,
				WhoseFault: req.WhoseFault,
				Status:     req.Status,
				Section:    req.Section,
				DisputedBy: employee.Name + employee.Surname,
				Company:    req.Company,
				LoadId:     logistic.Cargo.CargoID,
			}, tx)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
