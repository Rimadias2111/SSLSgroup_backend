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
// @Router /v1/transactions [post]
// @Summary Create a transaction
// @Description API for creating a new transaction
// @Tags transaction
// @Accept json
// @Produce json
// @Param transaction body swag.CreateUpdateTransaction true "Transaction data"
// @Success 200 {object} models.ResponseId
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) CreateTransaction(c *gin.Context) {
	var transactionModel swag.CreateUpdateTransaction
	if err := c.ShouldBindJSON(&transactionModel); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	puTime, err := time.Parse("2006-01-02T15:04:05", transactionModel.PuTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid pickup time format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	deliveryTime, err := time.Parse("2006-01-02T15:04:05", transactionModel.DeliveryTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid delivery time format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	driverId, err := uuid.Parse(transactionModel.DriverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid driver ID: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	employeeId, err := uuid.Parse(transactionModel.EmployeeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid employee ID: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	transaction := models.Transaction{
		From:         transactionModel.From,
		To:           transactionModel.To,
		PuTime:       puTime,
		DeliveryTime: deliveryTime,
		Success:      transactionModel.Success,
		LoadedMiles:  transactionModel.LoadedMiles,
		TotalMiles:   transactionModel.TotalMiles,
		Provider:     transactionModel.Provider,
		Cost:         transactionModel.Cost,
		Rate:         transactionModel.Rate,
		DriverId:     driverId,
		EmployeeId:   employeeId,
		CargoID:      transactionModel.CargoID,
	}

	id, err := h.service.Transaction().Create(c.Request.Context(), &transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while creating a transaction: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseId{Id: id})
}

// @Security ApiKeyAuth
// @Router /v1/transactions/{transaction_id} [put]
// @Summary Update a transaction
// @Description API for updating a transaction
// @Tags transaction
// @Accept json
// @Produce json
// @Param transaction_id path string true "Transaction ID"
// @Param transaction body swag.CreateUpdateTransaction true "Transaction data"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) UpdateTransaction(c *gin.Context) {
	var transactionModel swag.CreateUpdateTransaction
	transactionIdStr := c.Param("transaction_id")

	transactionId, err := uuid.Parse(transactionIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid transaction ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	if err := c.ShouldBindJSON(&transactionModel); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	puTime, err := time.Parse(time.RFC3339, transactionModel.PuTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid pickup time format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	deliveryTime, err := time.Parse(time.RFC3339, transactionModel.DeliveryTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid delivery time format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	driverId, err := uuid.Parse(transactionModel.DriverId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid driver ID: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	employeeId, err := uuid.Parse(transactionModel.EmployeeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid employee ID: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	transaction := models.Transaction{
		Id:           transactionId,
		From:         transactionModel.From,
		To:           transactionModel.To,
		PuTime:       puTime,
		DeliveryTime: deliveryTime,
		Success:      transactionModel.Success,
		LoadedMiles:  transactionModel.LoadedMiles,
		TotalMiles:   transactionModel.TotalMiles,
		Provider:     transactionModel.Provider,
		Cost:         transactionModel.Cost,
		Rate:         transactionModel.Rate,
		DriverId:     driverId,
		EmployeeId:   employeeId,
		CargoID:      transactionModel.CargoID,
	}

	if err := h.service.Transaction().Update(c.Request.Context(), &transaction); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while updating the transaction: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Transaction updated successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/transactions/{transaction_id} [delete]
// @Summary Delete a transaction
// @Description API for deleting a transaction
// @Tags transaction
// @Param transaction_id path string true "Transaction ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) DeleteTransaction(c *gin.Context) {
	transactionIdStr := c.Param("transaction_id")
	transactionId, err := uuid.Parse(transactionIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid transaction ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	err = h.service.Transaction().Delete(c.Request.Context(), models.RequestId{Id: transactionId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while deleting the transaction: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Transaction deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/transactions/{transaction_id} [get]
// @Summary Get a transaction by ID
// @Description API for retrieving a transaction by ID
// @Tags transaction
// @Param transaction_id path string true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetTransaction(c *gin.Context) {
	transactionIdStr := c.Param("transaction_id")
	transactionId, err := uuid.Parse(transactionIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid transaction ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	transaction, err := h.service.Transaction().Get(c.Request.Context(), models.RequestId{Id: transactionId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving the transaction: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// @Security ApiKeyAuth
// @Router /v1/transactions [get]
// @Summary Get all transactions
// @Description API for retrieving all transactions with pagination and search
// @Tags transaction
// @Param page query int false "Page number"
// @Param limit query int false "Number of transactions per page"
// @Param provider query string false "Service Provider"
// @Param success query bool false "Success"
// @Param cargo_id query string false "Cargo Id"
// @Param driver_name query string false "Driver Name"
// @Param dispatcher_name query string false "Dispatcher Name"
// @Success 200 {object} models.GetAllTransResp
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetAllTransactions(c *gin.Context) {
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

	cargoId := c.Query("cargo_id")

	provider := c.Query("provider")
	driverName := c.Query("driver_name")
	dispatcherName := c.Query("dispatcher_name")
	success := c.Query("success")
	if success != "" && success != "true" && success != "false" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid success: " + success,
			ErrorCode:    "Bad Request",
		})
		return
	}

	req := models.GetAllTransReq{
		Page:           page,
		Limit:          limit,
		Provider:       provider,
		CargoID:        cargoId,
		DriverName:     driverName,
		DispatcherName: dispatcherName,
		Success:        success,
	}

	transactions, err := h.service.Transaction().GetAll(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving transactions: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
