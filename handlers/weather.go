package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"weather_app/models"
	"weather_app/utils"
)

// WeatherHandler handles requests to /weather endpoint.
func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	weatherInfo := models.WeatherInfo{}
	lat, lon, units, err := utils.ValidateRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	config := utils.GetConfig()
	apiKey := config.APIKey

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&units=%s&appid=%s", lat, lon, units, apiKey)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching weather data:", err)
	}
	defer response.Body.Close()

	var weatherResp models.WeatherResponse
	err = json.NewDecoder(response.Body).Decode(&weatherResp)
	if err != nil {
		log.Fatal("Error decoding weather JSON:", err)
	}

	// Determine weather condition
	if len(weatherResp.Weather) > 0 {
		weatherCondition := weatherResp.Weather[0].Description
		// Determine temperature type
		var temperatureType string
		if units == "metric" {
			switch {
			case weatherResp.Main.Temp >= 30:
				temperatureType = "hot"
			case weatherResp.Main.Temp <= 10:
				temperatureType = "cold"
			default:
				temperatureType = "moderate"
			}
		} else if units == "imperial" {
			switch {
			case weatherResp.Main.Temp >= 86: // 30°C in Fahrenheit
				temperatureType = "hot"
			case weatherResp.Main.Temp <= 50: // 10°C in Fahrenheit
				temperatureType = "cold"
			default:
				temperatureType = "moderate"
			}
		}

		// Prepare response
		weatherInfo = models.WeatherInfo{
			Condition:       weatherCondition,
			TemperatureType: temperatureType,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherInfo)
}
