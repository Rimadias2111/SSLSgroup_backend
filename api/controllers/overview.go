package controllers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Security ApiKeyAuth
// @Router /v1/logistics/overview [get]
// @Summary Get driver overview
// @Description API to get an overview of driver statuses by company
// @Tags logistic
// @Accept json
// @Produce json
// @Success 200 {object} models.GetOverview "Overview data"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) Overview(c *gin.Context) {
	resp, err := h.service.Logistic().GetOverview(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: err.Error(),
			ErrorCode:    "InternalError",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
