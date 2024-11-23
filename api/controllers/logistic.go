package controllers

import (
	"backend/etc/Utime"
	"backend/models"
	"backend/models/swag"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

// @Security ApiKeyAuth
// @Router /v1/logistics [post]
// @Summary Create a logistic record
// @Description API for creating a new logistic record
// @Tags logistic
// @Accept json
// @Produce json
// @Param logistic body swag.CreateUpdateLogistic true "Logistic data"
// @Success 200 {object} models.ResponseId
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) CreateLogistic(c *gin.Context) {
	var logisticModel swag.CreateUpdateLogistic
	if err := c.ShouldBindJSON(&logisticModel); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	driverId, err := uuid.Parse(logisticModel.DriverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing driver id: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	cargoId, err := uuid.Parse(logisticModel.CargoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing cargo id: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	stTime, err := time.Parse("2006-01-02T15:04:05", logisticModel.StTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing start time: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	parts := strings.Split(logisticModel.Location, ",")
	if len(parts) != 2 {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing location: ",
			ErrorCode:    "Bad Request",
		})
		return
	}

	logistic := models.Logistic{
		DriverId:   driverId,
		CargoId:    &cargoId,
		Status:     logisticModel.Status,
		StTime:     &stTime,
		UpdateTime: Utime.Now(),
		Location:   logisticModel.Location,
		State:      strings.TrimSpace(parts[1]),
		Notion:     logisticModel.Notion,
		Post:       logisticModel.Post,
	}

	id, err := h.service.Logistic().Create(c.Request.Context(), &logistic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while creating a logistic record: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseId{Id: id})
}

// @Security ApiKeyAuth
// @Router /v1/logistics/{logistic_id} [put]
// @Summary Update a logistic record
// @Description API for updating a logistic record
// @Tags logistic
// @Accept json
// @Produce json
// @Param logistic_id path string true "Logistic ID"
// @Param logistic body swag.CreateUpdateLogistic true "Logistic data"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) UpdateLogistic(c *gin.Context) {
	var logisticModel swag.CreateUpdateLogistic
	logisticIdStr := c.Param("logistic_id")

	logisticId, err := uuid.Parse(logisticIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid logistic ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	if err := c.ShouldBindJSON(&logisticModel); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	cargoId, err := uuid.Parse(logisticModel.CargoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing cargo id: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	stTime, err := time.Parse("2006-01-02T15:04:05", logisticModel.StTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing start time: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	parts := strings.Split(logisticModel.Location, ",")
	if len(parts) != 2 {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing location: ",
			ErrorCode:    "Bad Request",
		})
		return
	}

	logistic := models.Logistic{
		Id:         logisticId,
		CargoId:    &cargoId,
		Status:     logisticModel.Status,
		StTime:     &stTime,
		UpdateTime: Utime.Now(),
		Location:   logisticModel.Location,
		State:      strings.TrimSpace(parts[1]),
		Notion:     logisticModel.Notion,
		Post:       logisticModel.Post,
	}

	if err := h.service.Logistic().Update(c.Request.Context(), &logistic); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while updating the logistic record: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Logistic record updated successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/logistics/{logistic_id} [delete]
// @Summary Delete a logistic record
// @Description API for deleting a logistic record
// @Tags logistic
// @Param logistic_id path string true "Logistic ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) DeleteLogistic(c *gin.Context) {
	logisticIdStr := c.Param("logistic_id")
	logisticId, err := uuid.Parse(logisticIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid logistic ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	err = h.service.Logistic().Delete(c.Request.Context(), models.RequestId{Id: logisticId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while deleting the logistic record: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Logistic record deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/logistics/{logistic_id} [get]
// @Summary Get a logistic record by ID
// @Description API for retrieving a logistic record by ID
// @Tags logistic
// @Param logistic_id path string true "Logistic ID"
// @Success 200 {object} models.Logistic
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetLogistic(c *gin.Context) {
	logisticIdStr := c.Param("logistic_id")
	logisticId, err := uuid.Parse(logisticIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing logistic ID: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	logistic, err := h.service.Logistic().Get(c.Request.Context(), models.RequestId{Id: logisticId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving the logistic record: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, logistic)
}

// @Security ApiKeyAuth
// @Router /v1/logistics [get]
// @Summary Get all logistic records
// @Description API for retrieving all logistic records with pagination and search
// @Tags logistic
// @Param page query int false "Page number"
// @Param limit query int false "Number of logistics per page"
// @Param driver_id query string false "Driver ID"
// @Success 200 {object} models.GetAllLogisticsResp
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetAllLogistics(c *gin.Context) {
	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid page: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid limit: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	req := models.GetAllLogisticsReq{
		Page:  page,
		Limit: limit,
	}

	logistics, err := h.service.Logistic().GetAll(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving logistics: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, logistics)
}

// @Security ApiKeyAuth
// @Router /v1/logistics_with_cargo/{logistic_id} [put]
// @Summary Update a logistic record with a cargo
// @Description API for updating a logistic record with a cargo
// @Tags logistic
// @Accept json
// @Produce json
// @Param logistic_id path string true "Logistic ID"
// @Param logistic body swag.UpdateLogisticWithCargo true "Logistic data"
// @Success 200 {object} models.ResponseId
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) UpdateLogisticCargo(c *gin.Context) {
	var logisticModel swag.UpdateLogisticWithCargo
	logisticIdStr := c.Param("logistic_id")

	err := c.ShouldBindJSON(&logisticModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing cargo id: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	logisticId, err := uuid.Parse(logisticIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid logistic ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	cargoId, err := uuid.Parse(logisticModel.CargoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing cargo id: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	stTime, err := time.Parse("2006-01-02T15:04:05", logisticModel.StTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing start time: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}
	var updateTime time.Time
	if logisticModel.Status == "COVERED" {
		updateTime, err = time.Parse("2006-01-02T15:04:05", logisticModel.PickUpTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseError{
				ErrorMessage: "Error while parsing pick up time: " + err.Error(),
				ErrorCode:    "Bad Request",
			})
			return
		}
	} else if logisticModel.Status == "ETA" || logisticModel.Status == "ETA WILL BE LATE" {
		updateTime, err = time.Parse("2006-01-02T15:04:05", logisticModel.DeliveryTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseError{
				ErrorMessage: "Error while parsing delivery time: " + err.Error(),
				ErrorCode:    "Bad Request",
			})
			return
		}
	} else {
		updateTime = Utime.Now()
	}

	parts := strings.Split(logisticModel.Location, ",")
	if len(parts) != 2 {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing location: ",
			ErrorCode:    "Bad Request",
		})
		return
	}

	logistic := models.Logistic{
		Id:         logisticId,
		CargoId:    &cargoId,
		Status:     logisticModel.Status,
		Post:       logisticModel.Post,
		Notion:     logisticModel.Notion,
		UpdateTime: updateTime,
		StTime:     &stTime,
		Location:   logisticModel.Location,
		State:      strings.TrimSpace(parts[1]),
	}

	pickUpTime, err := time.Parse("2006-01-02T15:04:05", logisticModel.PickUpTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing pick up time: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	deliveryTime, err := time.Parse("2006-01-02T15:04:05", logisticModel.DeliveryTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing delivery time: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	employeeId, err := uuid.Parse(logisticModel.EmployeeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing employee ID: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	if logisticModel.Provider == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing provider: ",
			ErrorCode:    "Bad Request",
		})
		return
	}

	cargo := models.Cargo{
		Id:           cargoId,
		PickUpTime:   pickUpTime,
		DeliveryTime: deliveryTime,
		Provider:     logisticModel.Provider,
		FreeMiles:    logisticModel.FreeMiles,
		LoadedMiles:  logisticModel.LoadedMiles,
		From:         logisticModel.From,
		To:           logisticModel.To,
		Cost:         logisticModel.Cost,
		Rate:         logisticModel.Rate,
		EmployeeId:   employeeId,
		CargoID:      logisticModel.LoadId,
	}

	if logisticModel.Status != "ETA, WILL BE LATE" {
		if stTime.After(deliveryTime) {
			c.JSON(http.StatusBadRequest, models.ResponseError{
				ErrorMessage: "Error delivery time cannot be less than ETA",
				ErrorCode:    "Bad Request",
			})
			return
		}
	}

	id, err := h.service.Logistic().UpdateWithCargo(c.Request.Context(), &logistic, &cargo, logisticModel.Create)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while updating logistic: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseId{
		Id: id,
	})
}

// @Security ApiKeyAuth
// @Router /v1/terminate_logistics [post]
// @Summary Terminate logistics
// @Description API for terminating logistic record
// @Tags logistic
// @Accept json
// @Produce json
// @Param logistic body swag.TerminateLogistic true "Logistic data"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) TerminateLogistic(c *gin.Context) {
	var req swag.TerminateLogistic

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	logisticId, err := uuid.Parse(req.LogisticId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	err = h.service.Logistic().Terminate(c.Request.Context(), models.RequestId{Id: logisticId}, req.Success)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Logistic Terminated",
	})
}

// @Security ApiKeyAuth
// @Router /v1/cancel_late_logistics [post]
// @Summary Cancel or late logistics
// @Description API for canceling or marking late logistic record
// @Tags logistic
// @Accept json
// @Produce json
// @Param logistic body swag.CancelLogistic true "Logistic data"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) CancelLateLogistic(c *gin.Context) {
	var req swag.CancelLogistic
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	logisticId, err := uuid.Parse(req.LogisticId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	if req.Company == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing company: ",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if req.Section == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing section: ",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if req.Status != "success" && req.Status != "canceled" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing status: ",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if req.WhoseFault != "Dispatcher" && req.WhoseFault != "Driver" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing whose fault: ",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if req.Reason == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing reason: ",
			ErrorCode:    "Bad Request",
		})
		return
	}

	empId, err := uuid.Parse("80c1d220-67b7-4dc9-82e4-393687b734a4")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	err = h.service.Logistic().CancelLate(c.Request.Context(), req, models.RequestId{Id: logisticId}, models.RequestId{Id: empId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Logistic Cancelled or made late",
	})
}
