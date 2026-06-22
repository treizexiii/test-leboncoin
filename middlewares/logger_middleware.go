package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		uuid := uuid.New()
		c.Set("request_id", uuid.String())

		log.Printf(
			"[%s] %s %s %s %s",
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Proto,
			uuid.String(),
		)

		// Process request
		c.Next()

		// Log details
		latency := time.Since(startTime)
		status := c.Writer.Status()

		log.Printf(
			"[%s] %s %s %s %s %d %v",
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Proto,
			uuid.String(),
			status,
			latency,
		)

	}
}
