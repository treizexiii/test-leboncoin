package handlers

import "github.com/gin-gonic/gin"

func RegisterRoutes(api *gin.RouterGroup) *gin.RouterGroup {
	api.GET("/fizzbuzz", FizzBuzzRoute)
	api.GET("/stats", StatsHandler)

	return api
}
