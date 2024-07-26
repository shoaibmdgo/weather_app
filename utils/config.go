package utils

import (
	"encoding/json"
	"log"
	"os"
)

// Config represents the structure of the configuration file.
type Config struct {
	APIKey string `json:"apiKey"`
}

// GetConfig reads the configuration file.
func GetConfig() Config {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal("Error opening config file:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error decoding config JSON:", err)
	}

	return config
}
