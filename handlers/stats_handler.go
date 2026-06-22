package handlers

import (
	"github.com/gin-gonic/gin"
)

func StatsHandler(c *gin.Context) {
	topRequests, count := counter.MostUsed()

	c.JSON(200, gin.H{
		"top_requests": topRequests,
		"count":        count,
	})
}
