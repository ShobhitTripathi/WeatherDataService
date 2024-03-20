package weather_api

import (
	"WeatherDataService/handler"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {
	router.GET("/weather/:zipcode/:days", handler.GetWeatherForecast)
	router.GET("/weather/:zipcode", handler.GetCurrentWeatherForecast)
}
