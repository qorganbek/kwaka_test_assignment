package service

import (
	"github.com/qorganbek/kwaka_test_assignment/internal/repository"
)

type WeatherService struct {
	repos repository.Weather
}

func NewWeatherService(repo repository.Weather) *WeatherService {
	return &WeatherService{repos: repo}
}
