package services

import (
	"backend/etc/Utime"
	"backend/models"
	"backend/models/swag"
	database "backend/st_database"
	"context"
	"errors"
	"fmt"
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

func (s *LogisticService) Update(ctx context.Context, req *models.Logistic, by models.RequestId) error {
	db := s.store.DB()
	err := db.Transaction(func(tx *gorm.DB) error {
		oldLogistic, getErr := s.store.Logistic().Get(ctx, models.RequestId{Id: req.Id})
		if getErr != nil {
			return getErr
		}

		if oldLogistic == nil {
			return fmt.Errorf("logistic with ID %s not found", req.Id)
		}

		err := s.store.Logistic().Update(ctx, req, tx)
		if err != nil {
			return err
		}

		_, err = s.store.History().Create(ctx, &models.History{
			DriverName: oldLogistic.Driver.Name + oldLogistic.Driver.Surname,
			LogisticId: req.Id,
			FromLogistic: models.JSONBLogistic{
				Post:       oldLogistic.Post,
				Status:     oldLogistic.Status,
				UpdateTime: oldLogistic.UpdateTime,
				StTime:     oldLogistic.StTime,
				State:      oldLogistic.State,
				Location:   oldLogistic.Location,
				Notion:     oldLogistic.Notion,
			},
			ToLogistic: models.JSONBLogistic{
				Post:       req.Post,
				Status:     req.Status,
				UpdateTime: req.UpdateTime,
				StTime:     req.StTime,
				State:      req.State,
				Location:   req.Location,
				Notion:     req.Notion,
			},
			FromCargo:  nil,
			ToCargo:    nil,
			EmployeeId: by.Id,
		}, tx)
		if err != nil {
			return err
		}

		return nil
	})
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

func (s *LogisticService) UpdateWithCargo(ctx context.Context, logistic *models.Logistic, cargo *models.Cargo, create bool, by models.RequestId) (string, error) {
	var (
		db  = s.store.DB()
		id  string
		err error
	)
	transErr := db.Transaction(func(tx *gorm.DB) error {
		oldLogistic, errG := s.store.Logistic().Get(ctx, models.RequestId{Id: logistic.Id})
		if errG != nil {
			return errG
		}

		if create && cargo.Id == uuid.Nil {
			id, err = s.store.Cargo().Create(ctx, cargo, tx)
			if err != nil {
				return err
			}

			cargoId, errP := uuid.Parse(id)
			if errP != nil {
				return errP
			}

			_, errH := s.store.History().Create(ctx, &models.History{
				DriverName: oldLogistic.Driver.Name + oldLogistic.Driver.Surname,
				LogisticId: logistic.Id,
				FromLogistic: models.JSONBLogistic{
					Post:       oldLogistic.Post,
					Status:     oldLogistic.Status,
					UpdateTime: oldLogistic.UpdateTime,
					StTime:     oldLogistic.StTime,
					State:      oldLogistic.State,
					Location:   oldLogistic.Location,
					Notion:     oldLogistic.Notion,
				},
				ToLogistic: models.JSONBLogistic{
					Post:       logistic.Post,
					Status:     logistic.Status,
					UpdateTime: logistic.UpdateTime,
					StTime:     logistic.StTime,
					State:      logistic.State,
					Location:   logistic.Location,
					Notion:     logistic.Notion,
				},
				FromCargo: nil,
				ToCargo: &models.JSONBCargo{
					Id:           cargo.Id,
					CargoID:      cargo.CargoID,
					Provider:     cargo.Provider,
					LoadedMiles:  cargo.LoadedMiles,
					FreeMiles:    cargo.FreeMiles,
					From:         cargo.From,
					To:           cargo.To,
					Cost:         cargo.Cost,
					Rate:         cargo.Rate,
					PickUpTime:   cargo.PickUpTime,
					DeliveryTime: cargo.DeliveryTime,
					EmployeeId:   cargo.EmployeeId,
				},
				EmployeeId: by.Id,
			})
			if errH != nil {
				return errH
			}

			logistic.CargoId = &cargoId
		} else {
			err = s.store.Cargo().Update(ctx, cargo, tx)
			if err != nil {
				return err
			}
			id = cargo.Id.String()

			oldCargo, errG := s.store.Cargo().Get(ctx, models.RequestId{Id: cargo.Id})
			if errG != nil {
				return errG
			}

			_, errH := s.store.History().Create(ctx, &models.History{
				DriverName: oldLogistic.Driver.Name + oldLogistic.Driver.Surname,
				LogisticId: logistic.Id,
				FromLogistic: models.JSONBLogistic{
					Post:       oldLogistic.Post,
					Status:     oldLogistic.Status,
					UpdateTime: oldLogistic.UpdateTime,
					StTime:     oldLogistic.StTime,
					State:      oldLogistic.State,
					Location:   oldLogistic.Location,
					Notion:     oldLogistic.Notion,
				},
				ToLogistic: models.JSONBLogistic{
					Post:       logistic.Post,
					Status:     logistic.Status,
					UpdateTime: logistic.UpdateTime,
					StTime:     logistic.StTime,
					State:      logistic.State,
					Location:   logistic.Location,
					Notion:     logistic.Notion,
				},
				FromCargo: &models.JSONBCargo{
					Id:           oldCargo.Id,
					CargoID:      oldCargo.CargoID,
					Provider:     oldCargo.Provider,
					LoadedMiles:  oldCargo.LoadedMiles,
					FreeMiles:    oldCargo.FreeMiles,
					From:         oldCargo.From,
					To:           oldCargo.To,
					Cost:         oldCargo.Cost,
					Rate:         oldCargo.Rate,
					PickUpTime:   oldCargo.PickUpTime,
					DeliveryTime: oldCargo.DeliveryTime,
					EmployeeId:   oldCargo.EmployeeId,
				},
				ToCargo: &models.JSONBCargo{
					Id:           cargo.Id,
					CargoID:      cargo.CargoID,
					Provider:     cargo.Provider,
					LoadedMiles:  cargo.LoadedMiles,
					FreeMiles:    cargo.FreeMiles,
					From:         cargo.From,
					To:           cargo.To,
					Cost:         cargo.Cost,
					Rate:         cargo.Rate,
					PickUpTime:   cargo.PickUpTime,
					DeliveryTime: cargo.DeliveryTime,
					EmployeeId:   cargo.EmployeeId,
				},
				EmployeeId: by.Id,
			})
			if errH != nil {
				return errH
			}
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

func (s *LogisticService) Terminate(ctx context.Context, req models.RequestId, success bool, by models.RequestId) error {
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

		if logistic.StTime == nil {
			time := Utime.Now()
			logistic.StTime = &time
		}

		errU := s.store.Logistic().Update(ctx, &models.Logistic{
			Id:         logistic.Id,
			Post:       false,
			Status:     "READY",
			UpdateTime: *logistic.StTime,
			StTime:     logistic.StTime,
			State:      logistic.State,
			Location:   logistic.Location,
			Emoji:      "",
			Notion:     "",
			CargoId:    nil,
		}, tx)
		if errU != nil {
			return errU
		}

		_, errH := s.store.History().Create(ctx, &models.History{
			DriverName: logistic.Driver.Name + logistic.Driver.Surname,
			LogisticId: logistic.Id,
			FromLogistic: models.JSONBLogistic{
				Post:       logistic.Post,
				Status:     logistic.Status,
				UpdateTime: logistic.UpdateTime,
				StTime:     logistic.StTime,
				State:      logistic.State,
				Location:   logistic.Location,
				Notion:     logistic.Notion,
			},
			ToLogistic: models.JSONBLogistic{
				Post:       false,
				Status:     "READY",
				UpdateTime: Utime.Now(),
				StTime:     logistic.StTime,
				State:      logistic.State,
				Location:   logistic.Location,
				Notion:     "",
			},
			FromCargo: &models.JSONBCargo{
				Id:           logistic.Cargo.Id,
				CargoID:      logistic.Cargo.CargoID,
				Provider:     logistic.Cargo.Provider,
				LoadedMiles:  logistic.Cargo.LoadedMiles,
				FreeMiles:    logistic.Cargo.FreeMiles,
				From:         logistic.Cargo.From,
				To:           logistic.Cargo.To,
				Cost:         logistic.Cargo.Cost,
				Rate:         logistic.Cargo.Rate,
				PickUpTime:   logistic.Cargo.PickUpTime,
				DeliveryTime: logistic.Cargo.DeliveryTime,
				EmployeeId:   logistic.Cargo.EmployeeId,
			},
			ToCargo:    nil,
			EmployeeId: by.Id,
		}, tx)
		if errH != nil {
			return errH
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *LogisticService) CancelLate(ctx context.Context, req swag.CancelLogistic, reqId models.RequestId, empId models.RequestId, compId models.RequestId) error {
	db := s.store.DB()

	err := db.Transaction(func(tx *gorm.DB) error {
		logistic, getErr := s.store.Logistic().Get(ctx, reqId)
		if getErr != nil {
			return getErr
		}

		if logistic.CargoId == nil {
			return errors.New("cargo not found")
		}

		if req.Cancel {
			_, err := s.store.Performance().Create(ctx, &models.Performance{
				Reason:     req.Reason,
				WhoseFault: req.WhoseFault,
				Status:     req.Status,
				Section:    req.Section,
				EmployeeId: empId.Id,
				CompanyId:  compId.Id,
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

			if logistic.StTime == nil {
				time := Utime.Now()
				logistic.StTime = &time
			}
			errU := s.store.Logistic().Update(ctx, &models.Logistic{
				Id:         logistic.Id,
				Post:       false,
				Status:     "READY",
				UpdateTime: *logistic.StTime,
				StTime:     logistic.StTime,
				State:      logistic.State,
				Location:   logistic.Location,
				Emoji:      "",
				Notion:     "",
				CargoId:    nil,
			}, tx)
			if errU != nil {
				return errU
			}

			_, errH := s.store.History().Create(ctx, &models.History{
				DriverName: logistic.Driver.Name + logistic.Driver.Surname,
				LogisticId: logistic.Id,
				FromLogistic: models.JSONBLogistic{
					Post:       logistic.Post,
					Status:     logistic.Status,
					UpdateTime: logistic.UpdateTime,
					StTime:     logistic.StTime,
					State:      logistic.State,
					Location:   logistic.Location,
					Notion:     logistic.Notion,
				},
				ToLogistic: models.JSONBLogistic{
					Post:       false,
					Status:     "READY",
					UpdateTime: Utime.Now(),
					StTime:     logistic.StTime,
					State:      logistic.State,
					Location:   logistic.Location,
					Notion:     "",
				},
				FromCargo: &models.JSONBCargo{
					Id:           logistic.Cargo.Id,
					CargoID:      logistic.Cargo.CargoID,
					Provider:     logistic.Cargo.Provider,
					LoadedMiles:  logistic.Cargo.LoadedMiles,
					FreeMiles:    logistic.Cargo.FreeMiles,
					From:         logistic.Cargo.From,
					To:           logistic.Cargo.To,
					Cost:         logistic.Cargo.Cost,
					Rate:         logistic.Cargo.Rate,
					PickUpTime:   logistic.Cargo.PickUpTime,
					DeliveryTime: logistic.Cargo.DeliveryTime,
					EmployeeId:   logistic.Cargo.EmployeeId,
				},
				ToCargo:    nil,
				EmployeeId: empId.Id,
			}, tx)
			if errH != nil {
				return errH
			}
		} else {
			_, err := s.store.Performance().Create(ctx, &models.Performance{
				Reason:     req.Reason,
				WhoseFault: req.WhoseFault,
				Status:     req.Status,
				Section:    req.Section,
				EmployeeId: empId.Id,
				CompanyId:  compId.Id,
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

func (s *LogisticService) GetOverview(ctx context.Context) (models.GetOverview, error) {
	resp, err := s.store.Logistic().Overview(ctx)
	if err != nil {
		return models.GetOverview{}, err
	}

	companies, err := s.store.Company().GetAll(ctx, models.GetAllCompaniesReq{Page: 1, Limit: 50})
	if err != nil {
		return models.GetOverview{}, err
	}

	CompanyMap := make(map[uuid.UUID]string)
	for _, company := range companies.Companies {
		CompanyMap[company.Id] = company.Name
	}

	for i := 0; i < len(resp.Companies); i++ {
		resp.Companies[i].Name = CompanyMap[resp.Companies[i].Id]
	}

	return resp, nil
}
