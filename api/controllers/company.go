package controllers

import (
	"backend/etc/filters"
	"backend/models"
	"backend/models/swag"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// @Security ApiKeyAuth
// @Router /v1/companies [post]
// @Summary Create a company
// @Description API for creating a new company
// @Tags company
// @Accept json
// @Produce json
// @Param company body swag.CreateUpdateCompany true "Company data"
// @Success 200 {object} models.ResponseId
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) CreateCompany(c *gin.Context) {
	var companyModel swag.CreateUpdateCompany
	if err := c.ShouldBindJSON(&companyModel); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	startDate, err := time.Parse("2006-01-02", companyModel.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid Start date format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	if companyModel.Name == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Name can't be empty",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if companyModel.DOT <= 0 {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Dot must be greater than zero",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if companyModel.SCAC == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "SAC can't be empty",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if !filters.ValidatePhoneNumber(companyModel.Number) {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid Number",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if companyModel.MC <= 0 {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Mc can't be zero",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if companyModel.Address == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Address can't be empty",
			ErrorCode:    "Bad Request",
		})
		return
	}

	company := models.Company{
		Name:      companyModel.Name,
		Address:   companyModel.Address,
		Number:    companyModel.Number,
		SCAC:      companyModel.SCAC,
		DOT:       companyModel.DOT,
		MC:        companyModel.MC,
		StartDate: &startDate,
	}

	id, err := h.service.Company().Create(c.Request.Context(), &company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while creating a company: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseId{Id: id})
}

// @Security ApiKeyAuth
// @Router /v1/companies/{company_id} [put]
// @Summary Update a company
// @Description API for updating a company
// @Tags company
// @Accept json
// @Produce json
// @Param company_id path string true "Company ID"
// @Param company body swag.CreateUpdateCompany true "Company data"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) UpdateCompany(c *gin.Context) {
	var companyModel swag.CreateUpdateCompany
	companyIdStr := c.Param("company_id")

	companyId, err := uuid.Parse(companyIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid company ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	if err := c.ShouldBindJSON(&companyModel); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	if companyModel.Name == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Name can't be empty",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if companyModel.DOT <= 0 {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Dot must be greater than zero",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if companyModel.SCAC == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "SAC can't be empty",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if !filters.ValidatePhoneNumber(companyModel.Number) {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid Number",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if companyModel.MC <= 0 {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Mc can't be zero",
			ErrorCode:    "Bad Request",
		})
		return
	}

	if companyModel.Address == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Address can't be empty",
			ErrorCode:    "Bad Request",
		})
		return
	}

	company := models.Company{
		Id:      companyId,
		Name:    companyModel.Name,
		Address: companyModel.Address,
		Number:  companyModel.Number,
		SCAC:    companyModel.SCAC,
		DOT:     companyModel.DOT,
		MC:      companyModel.MC,
	}

	if err := h.service.Company().Update(c.Request.Context(), &company); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while updating the company: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Company updated successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/companies/{company_id} [delete]
// @Summary Delete a company
// @Description API for deleting a company
// @Tags company
// @Param company_id path string true "Company ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) DeleteCompany(c *gin.Context) {
	idStr := c.Param("company_id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid company ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	err = h.service.Company().Delete(c.Request.Context(), models.RequestId{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while deleting the company: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Company deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/companies/{company_id} [get]
// @Summary Get a company by ID
// @Description API for retrieving a company by ID
// @Tags company
// @Param company_id path string true "Company ID"
// @Success 200 {object} models.Company
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetCompany(c *gin.Context) {
	idStr := c.Param("company_id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid company ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	company, err := h.service.Company().Get(c.Request.Context(), models.RequestId{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving the company: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, company)
}

// @Security ApiKeyAuth
// @Router /v1/companies [get]
// @Summary Get all companies
// @Description API for retrieving all companies with pagination and search
// @Tags company
// @Param page query int false "Page number"
// @Param limit query int false "Number of companies per page"
// @Param search query string false "Search term"
// @Success 200 {object} models.GetAllCompaniesResp
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetAllCompanies(c *gin.Context) {
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

	req := models.GetAllCompaniesReq{
		Page:  page,
		Limit: limit,
	}

	companies, err := h.service.Company().GetAll(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving companies: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, companies)
}
