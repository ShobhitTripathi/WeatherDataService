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
	"strconv"
	"time"
)

func GetWeatherForecast(ctx *gin.Context) {
	zipcode := ctx.Param("zipcode")
	days := ctx.Param("days")
	daysInt, err := strconv.Atoi(days)
	daysInt += constant.INT_ONE
	var response model.Response
	redis_key := zipcode + constant.UNDERSCORE + days

	// Validate input
	ok := validation.Validate(zipcode, daysInt)
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
	if entry, err := redisClient.Get(ctx, redis_key); entry != nil && err == nil {
		response.Cached = true
		response.WeatherData = entry.Data
		ctx.JSON(http.StatusOK, gin.H{"response": response})
		return
	}

	// If not in cache or expired, fetch data from Weather API
	data, err := service.FetchFutureWeatherFromAPI(zipcode, daysInt)
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
	err = redisClient.Set(ctx.Request.Context(), redis_key, &cacheEntry, time.Minute*30) // Cache for 30 minutes
	if err != nil {
		log.Printf("Failed to cache data in Redis: %v", err)
	}

	response.Cached = false
	response.WeatherData = data

	ctx.JSON(http.StatusOK, gin.H{"response": response})
}
