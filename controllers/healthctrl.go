package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
//// @Summary Healthcheck endpoint for Shop API
//// @Description This endpoint exists solely for checking the active status of the application
//// @Description Any Http status other than 200 signifies that the application is down
//// @Accept json
//// @Produce json
//// @Success 200
//// @Router /health [get]

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Status is ok"})
}
