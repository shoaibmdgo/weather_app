package utils

import (
	"errors"
	"net/http"
)

// ValidateRequest checks if the required query parameters are present.
func ValidateRequest(r *http.Request) (string, string, string, error) {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")
	units := r.URL.Query().Get("units")

	if lat == "" || lon == "" || units == "" {
		return "", "", "", errors.New("lat, lon, and units are required query parameters")
	}

	if units != "metric" && units != "imperial" {
		return "", "", "", errors.New("units must be 'metric' or 'imperial'")
	}

	return lat, lon, units, nil
}
