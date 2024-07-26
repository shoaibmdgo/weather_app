package main

import (
	"log"
	"net/http"
	"weather_app/handlers"
)

func main() {
	http.HandleFunc("/weather", handlers.WeatherHandler)
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
