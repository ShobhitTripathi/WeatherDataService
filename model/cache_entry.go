package model

import "time"

// CacheEntry is a struct that holds the weather data and its expiration time
type CacheEntry struct {
	Data       WeatherData
	Expiration time.Time
}
