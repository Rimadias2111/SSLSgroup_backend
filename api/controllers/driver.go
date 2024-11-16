package controllers

import (
	"backend/models"
	"backend/models/swag"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// @Security ApiKeyAuth
// @Router /v1/drivers [post]
// @Summary Create a driver
// @Description API for creating a new driver
// @Tags driver
// @Accept json
// @Produce json
// @Param driver body swag.CreateUpdateDriver true "Driver data"
// @Success 200 {object} models.ResponseId
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) CreateDriver(c *gin.Context) {
	var driverModel swag.CreateUpdateDriver
	if err := c.ShouldBindJSON(&driverModel); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	companyId, err := uuid.Parse(driverModel.CompanyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid Company ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	bTime, err := time.Parse("2006-01-02", driverModel.Birthday)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid Birthday format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
	}

	startDate, err := time.Parse("2006-01-02", driverModel.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid Start Date format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
	}

	if driverModel.Name == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Name field is required",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if driverModel.Surname == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Surname field is required",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if driverModel.TruckNumber == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "TruckNumber field is required",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if driverModel.Mail == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Mail field is required",
			ErrorCode:    "Bad Request",
		})
		return
	}

	driver := models.Driver{
		Name:        driverModel.Name,
		Surname:     driverModel.Surname,
		TruckNumber: driverModel.TruckNumber,
		PhoneNumber: driverModel.PhoneNumber,
		Mail:        driverModel.Mail,
		Birthday:    bTime,
		CompanyId:   companyId,
		StartDate:   &startDate,
	}

	id, err := h.service.Driver().Create(c.Request.Context(), &driver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while creating a driver: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseId{Id: id})
}

// @Security ApiKeyAuth
// @Router /v1/drivers/{driver_id} [put]
// @Summary Update a driver
// @Description API for updating a driver
// @Tags driver
// @Accept json
// @Produce json
// @Param driver_id path string true "Driver ID"
// @Param driver body swag.CreateUpdateDriver true "Driver data"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) UpdateDriver(c *gin.Context) {
	var driverModel swag.CreateUpdateDriver
	driverIdStr := c.Param("driver_id")

	driverId, err := uuid.Parse(driverIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid driver ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	if err := c.ShouldBindJSON(&driverModel); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	companyId, err := uuid.Parse(driverModel.CompanyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid Company ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	bTime, err := time.Parse("2006-01-02", driverModel.Birthday)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid Birthday format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
	}

	if driverModel.Name == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Name field is required",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if driverModel.Surname == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Surname field is required",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if driverModel.TruckNumber == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "TruckNumber field is required",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if driverModel.Mail == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Mail field is required",
			ErrorCode:    "Bad Request",
		})
		return
	}

	driver := models.Driver{
		Id:          driverId,
		Name:        driverModel.Name,
		Surname:     driverModel.Surname,
		TruckNumber: driverModel.TruckNumber,
		PhoneNumber: driverModel.PhoneNumber,
		Mail:        driverModel.Mail,
		Birthday:    bTime,
		CompanyId:   companyId,
	}

	if err := h.service.Driver().Update(c.Request.Context(), &driver); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while updating the driver: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Driver updated successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/drivers/{driver_id} [delete]
// @Summary Delete a driver
// @Description API for deleting a driver
// @Tags driver
// @Param driver_id path string true "Driver ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) DeleteDriver(c *gin.Context) {
	idStr := c.Param("driver_id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid driver ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	err = h.service.Driver().Delete(c.Request.Context(), models.RequestId{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while deleting the driver: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Driver deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/drivers/{driver_id} [get]
// @Summary Get a driver by ID
// @Description API for retrieving a driver by ID
// @Tags driver
// @Param driver_id path string true "Driver ID"
// @Success 200 {object} models.Driver
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetDriver(c *gin.Context) {
	idStr := c.Param("driver_id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid driver ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	driver, err := h.service.Driver().Get(c.Request.Context(), models.RequestId{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving the driver: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, driver)
}

// @Security ApiKeyAuth
// @Router /v1/drivers [get]
// @Summary Get all drivers
// @Description API for retrieving all drivers with pagination and search
// @Tags driver
// @Param page query int false "Page number"
// @Param limit query int false "Number of drivers per page"
// @Param name query string fasle "Drivers Name"
// @Param truck_number query int false "Truck Number"
// @Param search query string false "Search term"
// @Success 200 {object} models.GetAllDriversResp
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetAllDrivers(c *gin.Context) {
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

	truckNumber, err := ParseIntegerQueryParam(c, "truck_number")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid truck_number: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
	}

	name := c.Query("name")

	req := models.GetAllDriversReq{
		Page:        page,
		Limit:       limit,
		TruckNumber: truckNumber,
		Name:        name,
	}

	drivers, err := h.service.Driver().GetAll(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving drivers: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, drivers)
}
