package handler

import (
	"WeatherDataService/cache"
	"WeatherDataService/constant"
	"WeatherDataService/model"
	"WeatherDataService/service"
	"WeatherDataService/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func GetCurrentWeatherForecast(ctx *gin.Context) {
	zipcode := ctx.Param("zipcode")
	var response model.Response

	// Validate input
	ok := validation.Validate(zipcode, constant.INT_ZERO)
	if !ok {
		er := model.Error{
			Code: http.StatusBadRequest,
			Msg:  "Please enter a valid zip code.",
			Type: constant.VALIDATION_ERROR,
		}
		response.Error = er
		ctx.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	// Get RedisClient instance from context
	redisClient := ctx.Request.Context().Value(constant.REDIS_CLIENT_CONTEXT_KEY).(*cache.RedisCache)
	if entry, err := redisClient.Get(ctx, zipcode); entry != nil && err == nil {
		response.Cached = true
		response.WeatherData = entry.Data
		ctx.JSON(http.StatusOK, gin.H{"response": response})
		return
	}

	// If not in cache or expired, fetch data from Weather API
	data, err := service.FetchCurrentWeatherFromAPI(zipcode)
	if err != nil {
		er := model.Error{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
			Type: constant.VENDOR_ERROR,
		}
		response.Error = er
		ctx.JSON(http.StatusInternalServerError, gin.H{"response": response})
		return
	}

	// set data o Redis Cache
	cacheEntry := model.CacheEntry{
		Data:       data,
		Expiration: time.Now().Add(30 * time.Minute),
	}
	err = redisClient.Set(ctx.Request.Context(), zipcode, &cacheEntry, time.Minute*30) // Cache for 30 minutes
	if err != nil {
		log.Printf("Failed to cache data in Redis: %v", err)
	}

	response.Cached = false
	response.WeatherData = data

	ctx.JSON(http.StatusOK, gin.H{"response": response})
}
