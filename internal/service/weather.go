package service

import "kwaka_test/internal/entity"

func (w WeatherService) GetWeather(location string) (entity.Weather, error) {
	return w.repos.GetWeather(location)
}

func (w WeatherService) UpdateWeather(location string, input entity.UpdateWeather) error {
	return w.repos.UpdateWeather(location, input)
}

func (w WeatherService) CreateWeather(location string) (entity.Weather, error) {
	return w.repos.CreateWeather(location)
}
