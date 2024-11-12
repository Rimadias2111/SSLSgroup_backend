package controllers

import (
	"backend/models"
	"backend/models/swag"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// @Security ApiKeyAuth
// @Router /v1/employees [post]
// @Summary Create an employee
// @Description API for creating a new employee
// @Tags employee
// @Accept json
// @Produce json
// @Param employee body swag.CreateUpdateEmployee true "Employee data"
// @Success 200 {object} models.ResponseId
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) CreateEmployee(c *gin.Context) {
	var employeeModel swag.CreateUpdateEmployee
	if err := c.ShouldBindJSON(&employeeModel); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	logoId, err := uuid.Parse(employeeModel.LogoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid Logo ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	employee := models.Employee{
		Id:          uuid.New(),
		Name:        employeeModel.Name,
		Surname:     employeeModel.Surname,
		Username:    employeeModel.Username,
		Password:    employeeModel.Password,
		LogoId:      logoId,
		Email:       employeeModel.Email,
		PhoneNumber: employeeModel.PhoneNumber,
	}

	id, err := h.service.Employee().Create(c.Request.Context(), &employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while creating an employee: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseId{Id: id})
}

// @Security ApiKeyAuth
// @Router /v1/employees/{employee_id} [put]
// @Summary Update an employee
// @Description API for updating an employee
// @Tags employee
// @Accept json
// @Produce json
// @Param employee_id path string true "Employee ID"
// @Param employee body swag.CreateUpdateEmployee true "Employee data"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) UpdateEmployee(c *gin.Context) {
	var employeeModel swag.CreateUpdateEmployee
	employeeIdStr := c.Param("employee_id")

	employeeId, err := uuid.Parse(employeeIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid employee ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	if err := c.ShouldBindJSON(&employeeModel); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	logoId, err := uuid.Parse(employeeModel.LogoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid Logo ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	employee := models.Employee{
		Id:          employeeId,
		Name:        employeeModel.Name,
		Surname:     employeeModel.Surname,
		Username:    employeeModel.Username,
		Password:    employeeModel.Password,
		LogoId:      logoId,
		Email:       employeeModel.Email,
		PhoneNumber: employeeModel.PhoneNumber,
	}

	if err := h.service.Employee().Update(c.Request.Context(), &employee); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while updating the employee: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Employee updated successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/employees/{employee_id} [delete]
// @Summary Delete an employee
// @Description API for deleting an employee
// @Tags employee
// @Param employee_id path string true "Employee ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) DeleteEmployee(c *gin.Context) {
	employeeIdStr := c.Param("employee_id")
	employeeId := uuid.MustParse(employeeIdStr)

	err := h.service.Employee().Delete(c.Request.Context(), models.RequestId{Id: employeeId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while deleting the employee: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Employee deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/employees/{employee_id} [get]
// @Summary Get an employee by ID
// @Description API for retrieving an employee by ID
// @Tags employee
// @Param employee_id path string true "Employee ID"
// @Success 200 {object} models.Employee
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetEmployee(c *gin.Context) {
	employeeIdStr := c.Param("employee_id")
	employeeId := uuid.MustParse(employeeIdStr)

	employee, err := h.service.Employee().Get(c.Request.Context(), models.RequestId{Id: employeeId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving the employee: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// @Security ApiKeyAuth
// @Router /v1/employees [get]
// @Summary Get all employees
// @Description API for retrieving all employees with pagination and search
// @Tags employee
// @Param page query int false "Page number"
// @Param limit query int false "Number of employees per page"
// @Param name query string false "Employee Name"
// @Param username query string false "Employee Username"
// @Param search query string false "Search term"
// @Success 200 {object} models.GetAllEmployeesResp
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetAllEmployees(c *gin.Context) {
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

	name := c.Query("name")
	username := c.Query("username")

	req := models.GetAllEmployeesReq{
		Page:     page,
		Limit:    limit,
		Name:     name,
		Username: username,
	}

	employees, err := h.service.Employee().GetAll(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving employees: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, employees)
}

