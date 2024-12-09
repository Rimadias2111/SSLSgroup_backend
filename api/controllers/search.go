package controllers

import (
	"backend/etc/search"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Controller) SearchHandler(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр 'q' не указан"})
		return
	}

	results, err := search.GetLocations(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}
