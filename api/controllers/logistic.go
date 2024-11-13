package controllers

import (
	"backend/etc/Utime"
	"backend/models"
	"backend/models/swag"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
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
	}

	cargoId, err := uuid.Parse(logisticModel.CargoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing cargo id: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
	}

	stTime, err := time.Parse("2006-01-02T15:04:05Z07:00", logisticModel.StTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing start time: " + err.Error(),
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
		Notion:     logisticModel.Notion,
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

	driverId, err := uuid.Parse(logisticModel.DriverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing driver id: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
	}

	cargoId, err := uuid.Parse(logisticModel.CargoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing cargo id: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
	}

	stTime, err := time.Parse("2006-01-02T15:04:05Z07:00", logisticModel.StTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing start time: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	logistic := models.Logistic{
		Id:         logisticId,
		DriverId:   driverId,
		CargoId:    &cargoId,
		Status:     logisticModel.Status,
		StTime:     &stTime,
		UpdateTime: Utime.Now(),
		Location:   logisticModel.Location,
		Notion:     logisticModel.Notion,
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
	logisticId := uuid.MustParse(logisticIdStr)

	err := h.service.Logistic().Delete(c.Request.Context(), models.RequestId{Id: logisticId})
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
	logisticId := uuid.MustParse(logisticIdStr)

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
