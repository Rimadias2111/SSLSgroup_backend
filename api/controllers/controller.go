package controllers

import (
	"backend/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Controller struct {
	service service.IService
}

func NewController(serviceS service.IService) *Controller {
	return &Controller{service: serviceS}
}

func ParsePageQueryParam(c *gin.Context) (uint64, error) {
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.ParseUint(pageStr, 10, 30)
	if err != nil {
		return 0, err
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

func ParseLimitQueryParam(c *gin.Context) (uint64, error) {
	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.ParseUint(limitStr, 10, 30)
	if err != nil {
		return 0, err
	}

	if limit == 0 {
		return 10, nil
	}
	return limit, nil
}

func ParseIntegerQueryParam(c *gin.Context, query string) (int64, error) {
	intStr := c.Query(query)
	if intStr == "" {
		return 0, nil
	}
	intA, err := strconv.ParseInt(intStr, 10, 30)
	if err != nil {
		return 0, err
	}

	return intA, nil
}
