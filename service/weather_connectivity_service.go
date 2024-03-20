package service

import (
	"WeatherDataService/constant"
	"WeatherDataService/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func FetchCurrentWeatherFromAPI(zipcode string) (model.WeatherData, error) {
	// Api key only available till 2nd April 2024
	apiKey := constant.API_KEY
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, zipcode)

	resp, err := http.Get(url)
	if err != nil {
		return model.WeatherData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.WeatherData{}, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// resolve the issue below
	weatherResponse := model.WeatherResponse{}
	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		return model.WeatherData{}, err
	}

	// Map WeatherResponse data to WeatherData struct
	currentTime := time.Now()
	weatherData := model.WeatherData{
		Zipcode: zipcode,
		City:    weatherResponse.Location.Name,
		Country: weatherResponse.Location.Country,
		CurrentWeather: &model.DayData{
			CurrentTempC:  weatherResponse.Current.TemperatureC,
			FeelsLikeC:    weatherResponse.Current.FeelslikeC,
			Humidity:      weatherResponse.Current.Humidity,
			Precipitation: weatherResponse.Current.PrecipMm,
			Forecast:      weatherResponse.Current.Condition.Text,
		},
		LastFetched: &currentTime,
	}

	return weatherData, nil
}

func FetchFutureWeatherFromAPI(zipcode string, days int) (model.WeatherData, error) {
	// Api key only available till 2nd April 2024
	apiKey := constant.API_KEY
	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&aqi=no&days=%d", apiKey, zipcode, days)

	resp, err := http.Get(url)
	if err != nil {
		return model.WeatherData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.WeatherData{}, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	weatherResponse := model.WeatherAPIResponse{}
	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		return model.WeatherData{}, err
	}

	currentTime := time.Now()
	weatherData := model.WeatherData{
		Zipcode:     zipcode,
		City:        weatherResponse.Location.Name,
		Country:     weatherResponse.Location.Country,
		LastFetched: &currentTime,
	}

	for _, forecastDay := range weatherResponse.Forecast.ForecastDay {
		var dayDataList []model.DayData
		forecastDate := time.Unix(int64(forecastDay.DateEpoch), 0)
		currentDate := time.Now()
		if forecastDate.Year() == currentDate.Year() && forecastDate.YearDay() == currentDate.YearDay() {
			// forecastDay.DateEpoch is today's date
			currentWeather := model.DayData{
				Day:           "Today",
				Date:          time.Unix(int64(forecastDay.DateEpoch), 0),
				CurrentTempC:  weatherResponse.Current.TemperatureC,
				MinTempC:      forecastDay.Day.MinTempC,
				MaxTempC:      forecastDay.Day.MaxTempC,
				FeelsLikeC:    weatherResponse.Current.FeelslikeC,
				Humidity:      weatherResponse.Current.Humidity,
				Precipitation: weatherResponse.Current.PrecipMm,
				Forecast:      weatherResponse.Current.Condition.Text,
			}
			weatherData.CurrentWeather = &currentWeather
		} else {
			// forecastDay.DateEpoch is not today's date
			dayDataList = append(dayDataList, model.DayData{
				Day:           forecastDate.Weekday().String(),
				Date:          time.Unix(int64(forecastDay.DateEpoch), 0),
				MaxTempC:      forecastDay.Day.MaxTempC,
				MinTempC:      forecastDay.Day.MinTempC,
				FeelsLikeC:    forecastDay.Day.AvgTempC,
				Humidity:      forecastDay.Day.AvgHumidity,
				Precipitation: forecastDay.Day.TotalPrecipMm,
				Forecast:      forecastDay.Day.Condition.Text,
			})
		}
		weatherData.DayData = append(weatherData.DayData, dayDataList...)
	}

	return weatherData, nil
}
