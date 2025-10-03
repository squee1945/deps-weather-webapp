package app

import (
	"fmt"
	ipweather "github.com/squee1945/deps-ip-weather"
)

// WeatherAdapter adapts the ipweather.IPWeather to the WeatherGetter interface.
type WeatherAdapter struct {
	ipw *ipweather.IPWeather
}

// NewWeatherAdapter creates a new WeatherAdapter.
func NewWeatherAdapter() (*WeatherAdapter, error) {
	ipw, err := ipweather.New()
	if err != nil {
		return nil, fmt.Errorf("ipweather.New: %w", err)
	}
	return &WeatherAdapter{ipw: ipw}, nil
}

// GetWeather implements the WeatherGetter interface.
func (a *WeatherAdapter) GetWeather(ip string) (string, error) {
	details, err := a.ipw.GetWeather(ip)
	if err != nil {
		return "", fmt.Errorf("ipw.GetWeather: %w", err)
	}
	return details.Conditions, nil
}
