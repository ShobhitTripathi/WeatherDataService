package main

import (
	"WeatherDataService/api"
	"WeatherDataService/cache"
	"WeatherDataService/constant"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Initialize Redis cache
	redisClient := cache.NewRedisCache(constant.LOCAL_HOST+constant.COLON+constant.REDIS_PORT, constant.EMPTY_STRING, constant.INT_ZERO)
	// Initialize Gin router
	router := gin.Default()

	// Middleware to inject RedisClient instance into context
	router.Use(func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), constant.REDIS_CLIENT_CONTEXT_KEY, redisClient)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	})

	// initialize api
	api.Init(&router.RouterGroup)
	router.SetTrustedProxies([]string{"127.0.0.1"})
	if err := router.Run(constant.COLON + constant.PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		log.Println("Application Running on port :", constant.PORT)
	}
	fmt.Printf("Application Running on port : %s\n", constant.PORT)
}
