package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/treizexiii/test-leboncoin/services"
)

type FizzBuzzRequest struct {
	Num1  int    `json:"num1"`
	Num2  int    `json:"num2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

func FizzBuzzRoute(c *gin.Context) {
	var request FizzBuzzRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := services.FizzBuzz(request.Num1, request.Num2, request.Limit, request.Str1, request.Str2)

	c.JSON(200, gin.H{
		"result": result,
	})
}
