package service

import (
	"kwaka_test/internal/repository"
)

type WeatherService struct {
	repos repository.Weather
}

func NewWeatherService(repo repository.Weather) *WeatherService {
	return &WeatherService{repos: repo}
}
