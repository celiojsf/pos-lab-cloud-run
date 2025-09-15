package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type WeatherService struct {
	client     *http.Client
	apiKey     string
	apiBaseURL string
}

type WeatherResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

type Temperature struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

func NewWeatherService() *WeatherService {
	return &WeatherService{
		client:     &http.Client{},
		apiKey:     os.Getenv("WEATHER_API_KEY"),
		apiBaseURL: "http://api.weatherapi.com/v1",
	}
}

func (s *WeatherService) GetTemperature(city string) (*Temperature, error) {
	url := fmt.Sprintf("%s/current.json?key=%s&q=%s", s.apiBaseURL, s.apiKey, city)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather API error: %d", resp.StatusCode)
	}

	var weatherData WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, err
	}

	celsius := weatherData.Current.TempC
	fahrenheit := celsius*1.8 + 32
	kelvin := celsius + 273.15

	return &Temperature{
		Celsius:    celsius,
		Fahrenheit: fahrenheit,
		Kelvin:     kelvin,
	}, nil
}
