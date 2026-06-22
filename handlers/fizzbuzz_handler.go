package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/treizexiii/test-leboncoin/services"
)

func FizzBuzzRoute(c *gin.Context) {
	var request *services.FizzBuzzRequest
	c.ShouldBindBodyWithJSON(&request)
	if request == nil || !request.Validate() {
		c.JSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}
	counter.Increment(*request)

	result := services.FizzBuzz(request.Num1, request.Num2, request.Limit, request.Str1, request.Str2)

	c.JSON(200, gin.H{
		"result": result,
	})
}
