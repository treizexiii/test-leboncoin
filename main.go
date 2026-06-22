package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/treizexiii/test-leboncoin/handlers"
	"github.com/treizexiii/test-leboncoin/middlewares"
)

func main() {
	if err := Load("."); err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	srv := gin.New()

	config := cors.Config{
		AllowOrigins:     Cfg.Cors.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	api := srv.Group("/api", cors.New(config))

	api.Use(middlewares.LoggerMiddleware())
	// api.Use(middlewares.DeserializeJSON(&services.FizzBuzzRequest{}))

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	handlers.RegisterRoutes(api)

	srv.Run(fmt.Sprintf(":%d", Cfg.Server.Port))
}
