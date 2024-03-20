package model

import (
	"time"
)

type Response struct {
	Cached      bool        `json:"cached"`
	WeatherData WeatherData `json:"weather_data,omitempty"`
	Error       Error       `json:"error,omitempty"`
}

// WeatherData is a struct that holds the weather data for a given location
type WeatherData struct {
	CurrentWeather *DayData   `json:"current_weather,omitempty"`
	DayData        []DayData  `json:"forecast_weather,omitempty"`
	Zipcode        string     `json:"zipcode,omitempty"`
	Country        string     `json:"country,omitempty"`
	City           string     `json:"city,omitempty"`
	LastFetched    *time.Time `json:"last_fetched,omitempty"`
}

type DayData struct {
	Day           string    `json:"day,omitempty"`
	Date          time.Time `json:"date,omitempty"`
	CurrentTempC  float64   `json:"current_temp_c,omitempty"`
	MaxTempC      float64   `json:"max_temp_c,omitempty"`
	MinTempC      float64   `json:"min_temp_c,omitempty"`
	FeelsLikeC    float64   `json:"feels_like_c,omitempty"`
	Humidity      int       `json:"humidity,omitempty"`
	Precipitation float64   `json:"precipitation,omitempty"`
	Forecast      string    `json:"forecast,omitempty"`
}
