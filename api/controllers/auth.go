package controllers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Security ApiKeyAuth
// @Router /v1/login [post]
// @Summary Log in
// @Description API for login
// @Tags auth
// @Accept json
// @Produce json
// @Param employee body models.AuthReq true "Employee data"
// @Success 200 {object} models.AuthResp
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) Login(c *gin.Context) {
	var req models.AuthReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "body did not contain required fields",
			ErrorCode:    "BAD_REQUEST",
		})
		return
	}

	resp, err := h.service.Employee().Auth(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: err.Error(),
			ErrorCode:    "SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
