package controllers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// @Security ApiKeyAuth
// @Router /v1/histories/{history_id} [get]
// @Summary Get a history record by ID
// @Description API for retrieving a single history record by ID
// @Tags history
// @Param history_id path string true "History ID"
// @Success 200 {object} models.History
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetHistory(c *gin.Context) {
	idStr := c.Param("history_id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid history ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	history, err := h.service.History().GetHistory(c.Request.Context(), models.RequestId{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving the history record: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, history)
}

// @Security ApiKeyAuth
// @Router /v1/histories [get]
// @Summary Get all history records
// @Description API for retrieving all history records with pagination and filters
// @Tags history
// @Param page query int false "Page number"
// @Param limit query int false "Number of records per page"
// @Success 200 {object} models.GetAllHistoriesResp
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetAllHistories(c *gin.Context) {
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

	req := models.GetAllHistoryReq{
		Page:  page,
		Limit: limit,
	}

	histories, err := h.service.History().GetAll(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving histories: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, histories)
}
