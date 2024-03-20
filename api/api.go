package api

import (
	"WeatherDataService/api/weather_api"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {
	router = router.Group("/api")
	weather_api.Init(router)
}
