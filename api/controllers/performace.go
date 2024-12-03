package controllers

import (
	"backend/models"
	"backend/models/swag"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// @Security ApiKeyAuth
// @Router /v1/performances [post]
// @Summary Create a performance
// @Description API for creating a new performance
// @Tags performance
// @Accept json
// @Produce json
// @Param performance body swag.CreateUpdatePerformance true "Performance data"
// @Success 200 {object} models.ResponseId
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) CreatePerformance(c *gin.Context) {
	var performanceModel swag.CreateUpdatePerformance
	if err := c.ShouldBindJSON(&performanceModel); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	employeeId, err := uuid.Parse(performanceModel.EmployeeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing employee id: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	companyId, err := uuid.Parse(performanceModel.CompanyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing company id: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	performance := models.Performance{
		Reason:     performanceModel.Reason,
		WhoseFault: performanceModel.WhoseFault,
		Status:     performanceModel.Status,
		LoadId:     performanceModel.LoadId,
		Section:    performanceModel.Section,
		EmployeeId: employeeId,
		CompanyId:  companyId,
	}

	id, err := h.service.Performance().Create(c.Request.Context(), &performance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while creating a performance: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseId{Id: id})
}

// @Security ApiKeyAuth
// @Router /v1/performances/{performance_id} [put]
// @Summary Update a performance
// @Description API for updating a performance
// @Tags performance
// @Accept json
// @Produce json
// @Param performance_id path string true "Performance ID"
// @Param performance body swag.CreateUpdatePerformance true "Performance data"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) UpdatePerformance(c *gin.Context) {
	var performanceModel swag.CreateUpdatePerformance
	performanceIdStr := c.Param("performance_id")

	performanceId, err := uuid.Parse(performanceIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid performance ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	if err := c.ShouldBindJSON(&performanceModel); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	employeeId, err := uuid.Parse(performanceModel.EmployeeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing employee id: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	companyId, err := uuid.Parse(performanceModel.CompanyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while parsing company id: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	performance := models.Performance{
		Id:         performanceId,
		Reason:     performanceModel.Reason,
		WhoseFault: performanceModel.WhoseFault,
		Status:     performanceModel.Status,
		LoadId:     performanceModel.LoadId,
		Section:    performanceModel.Section,
		EmployeeId: employeeId,
		CompanyId:  companyId,
	}

	if err := h.service.Performance().Update(c.Request.Context(), &performance); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while updating the performance: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Performance updated successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/performances/{performance_id} [delete]
// @Summary Delete a performance
// @Description API for deleting a performance
// @Tags performance
// @Param performance_id path string true "Performance ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) DeletePerformance(c *gin.Context) {
	performanceIdStr := c.Param("performance_id")
	performanceId, err := uuid.Parse(performanceIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid performance ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	err = h.service.Performance().Delete(c.Request.Context(), models.RequestId{Id: performanceId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while deleting the performance: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Performance deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/performances/{performance_id} [get]
// @Summary Get a performance by ID
// @Description API for retrieving a performance by ID
// @Tags performance
// @Param performance_id path string true "Performance ID"
// @Success 200 {object} models.Performance
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetPerformance(c *gin.Context) {
	performanceIdStr := c.Param("performance_id")
	performanceId, err := uuid.Parse(performanceIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid performance ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	performance, err := h.service.Performance().Get(c.Request.Context(), models.RequestId{Id: performanceId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving the performance: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, performance)
}

// @Security ApiKeyAuth
// @Router /v1/performances [get]
// @Summary Get all performances
// @Description API for retrieving all performances with pagination and search
// @Tags performance
// @Param page query int false "Page number"
// @Param limit query int false "Number of performances per page"
// @Param company query string false "Company"
// @Param whose_fault query string false "Whose fault"
// @Param status query string false "Status"
// @Param section query string false "Section"
// @Param disputed_by query string false "Disputed by"
// @Success 200 {object} models.GetAllPerformancesResp
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetAllPerformances(c *gin.Context) {
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

	companyIdStr := c.Query("company_id")
	var companyId = uuid.Nil
	if companyIdStr != "" {
		companyId, err = uuid.Parse(companyIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseError{
				ErrorMessage: "Invalid company ID format: " + err.Error(),
				ErrorCode:    "Bad Request",
			})
			return
		}
	}

	employeeIdStr := c.Query("employee_id")
	var employeeId = uuid.Nil
	if employeeIdStr != "" {
		employeeId, err = uuid.Parse(employeeIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseError{
				ErrorMessage: "Invalid employee ID format: " + err.Error(),
				ErrorCode:    "Bad Request",
			})
			return
		}
	}

	section := c.Query("section")
	whoseFault := c.Query("whose_fault")
	if whoseFault != "" && whoseFault != "Driver" && whoseFault != "Dispatcher" && whoseFault != "Company" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid query parameter: ",
			ErrorCode:    "Bad Request",
		})
		return
	}

	status := c.Query("status")
	if status != "" && status != "success" && status != "canceled" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid query parameter: ",
			ErrorCode:    "Bad Request",
		})
		return
	}
	req := models.GetAllPerformancesReq{
		Page:       page,
		Limit:      limit,
		CompanyId:  companyId,
		Section:    section,
		WhoseFault: whoseFault,
		Status:     status,
		EmployeeId: employeeId,
	}

	performances, err := h.service.Performance().GetAll(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving performances: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, performances)
}
