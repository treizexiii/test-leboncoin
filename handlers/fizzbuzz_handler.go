package handlers

import "github.com/gin-gonic/gin"

func FizzBuzzRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "FizzBuzz route",
	})
}
