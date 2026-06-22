package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/treizexiii/test-leboncoin/services"
)

// Registered services
var counter = services.NewRequestCounter()

func RegisterRoutes(api *gin.RouterGroup) *gin.RouterGroup {
	api.GET("/fizzbuzz", FizzBuzzRoute)
	api.GET("/stats", StatsHandler)

	return api
}
