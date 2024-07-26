package models

// WeatherResponse represents the structure of the response from OpenWeatherMap API.
type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

// WeatherInfo represents the structure of the response from our server.
type WeatherInfo struct {
	Condition       string `json:"condition"`
	TemperatureType string `json:"temperatureType"`
}
